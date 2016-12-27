#!/usr/bin/python3.6
import os
import sys
import random
import base64
import glob
import datetime
from colorama import init
from termcolor import colored
import threading
import time
import vlc
import flask
import configparser
from peewee import *
from playhouse.sqlite_ext import SqliteExtDatabase

DB = SqliteExtDatabase('main.sqlite')


class BaseModel(Model):

    class Meta:
        database = DB


class Alarm(BaseModel):
    name = CharField()
    source = CharField()
    url = CharField()
    alarmTime = CharField()
    created = DateTimeField(default=datetime.datetime.now)
    lastAlarmDate = CharField(default="")
    repeatDays = CharField()
    active = BooleanField(default=True)
    playing = BooleanField(default=False)


class MainMusicAlarm:
    Config = configparser.ConfigParser()
    PlayingTitle = vlc.MediaPlayer()
    ConfigFileName = 'config.ini'

    def __init__(self):
        self.Config.read(self.ConfigFileName)
        print(colored('Initializing the music Alarm checker at {}.'.format(
            datetime.datetime.now().strftime("%H:%M")), 'green'))
        self.initDB()

    def parseTime(self, time):
        h, m, s = map(int, time.split(':'))
        return datetime.datetime.now().replace(hour=h, minute=m, second=s, microsecond=0)

    def playFile(self, alarm, state):
        if state:
            self.PlayingTitle = vlc.MediaPlayer(alarm.url)
            self.PlayingTitle.play()
            alarm.lastAlarmDate = datetime.datetime.now().strftime("%d.%m.%Y")
            alarm.playing = True
        else:
            self.PlayingTitle.stop()
            alarm.playing = False
        alarm.save()

    def checkAlarms(self):
        # print(datetime.datetime.now().strftime("%H:%M"))
        for alarm in Alarm.select():
            # check if the alarm is over the current time
            if self.parseTime(alarm.alarmTime) < datetime.datetime.now() and self.parseTime(alarm.alarmTime) > alarm.created:
                # check if the alarm is currently not playing
                if alarm.playing == False and alarm.active == True:
                    # check if the alarm was not already played on this day
                    if str(self.parseTime(alarm.alarmTime).weekday() + 1) in alarm.repeatDays.strip().split(','):
                        if alarm.lastAlarmDate != datetime.datetime.now().strftime("%d.%m.%Y"):
                            print(colored('Triggering the alarm {} at {} for {}.'.format(
                                alarm.name, datetime.datetime.now(), self.parseTime(alarm.alarmTime)), 'green'))
                            self.manageAlarmTargets(alarm)
        time.sleep(10)
        self.checkAlarms()

    def manageAlarmTargets(self, alarm):
        if alarm.source == 'file':
            self.playFile(alarm, 1)
        os.system("sudo send433 10101 4 {}".format(1))

    def stopAlarm(self):
        for alarm in Alarm.select().where(Alarm.playing == 1):
            alarm.playing = False
            alarm.save()
            os.system("sudo send433 10101 4 {}".format(0))

    def getFiles(self):
        pass

    def initDB(self):
        DB.connect()
        Alarm.create_table(True)


app = flask.Flask(__name__)
M = MainMusicAlarm()


@app.route('/api/v1/stopAlarms')
def stopAlarm():
    M.stopAlarm()
    return flask.jsonify({'success': True})


@app.route('/api/v1/createAlarm', methods=['POST'])
def createAlarm_page():
    Alarm.create(name=flask.request.json['name'],
                 source=flask.request.json['source'],
                 url=flask.request.json['url'],
                 alarmTime=flask.request.json['alarmTime'],
                 repeatDays=flask.request.json['repeatDays'],
                 active=flask.request.json['active'])
    return flask.jsonify({'success': True})


@app.route('/api/v1/deleteAlarm/<int:AlarmId>')
def deleteAlarm_page(AlarmId):
    return flask.jsonify(Alarm.get(Alarm.id == AlarmId).delete_instance())


@app.route('/api/v1/modifyAlarm/<int:AlarmId>', methods=['POST'])
def modifyAlarm_page(AlarmId):
    entry = Alarm.get(Alarm.id == AlarmId)
    entry.name = flask.request.json[
        'name'] if 'name' in flask.request.json else entry.name
    entry.source = flask.request.json[
        'source'] if 'source' in flask.request.json else entry.name
    entry.url = flask.request.json[
        'url'] if 'url' in flask.request.json else entry.name
    entry.alarmTime = flask.request.json[
        'alarmTime'] if 'alarmTime' in flask.request.json else entry.name
    entry.repeatDays = flask.request.json[
        'repeatDays'] if 'repeatDays' in flask.request.json else entry.name
    entry.active = flask.request.json[
        'active'] if 'active' in flask.request.json else entry.name
    entry.save()
    return flask.jsonify({'success': True})


@app.route('/api/v1/getAlarms')
def getAlarms():
    return flask.jsonify(list(Alarm.select().dicts()))


@app.route('/<path:path>')
def default_serve(path):
    return flask.send_from_directory('static', path)


@app.route('/api/v1/getFiles')
def getFiles_page():
    files = []
    for filename in glob.iglob(M.Config['ALARM']['musicdirectory'], recursive=True):
        files.append(filename)
    return flask.jsonify(files)


@app.route('/')
def main_page():
    return flask.render_template('index.html', site="Python Application Webinterface")


@app.route('/alarms')
def alarms_page():
    return flask.render_template('alarms.html', site="Alarms", alarms=Alarm.select().dicts())


@app.route('/history')
def history_page():
    return flask.render_template('history.html', site="History")

if (__name__ == "__main__"):
    threading.Thread(target=M.checkAlarms).start()
    app.run(host=M.Config['HTTP']['ListenAddr'], port=M.Config[
            'HTTP']['ListenPort'], threaded=True)
    print(colored('The programm was terminated at {}. Stopping services...'.format(
        datetime.datetime.now()), 'red'))
    M.manageAlarm(0)
    sys.exit(os.EX_OK)
