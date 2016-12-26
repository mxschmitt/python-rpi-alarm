<<<<<<< HEAD
#!/usr/bin/python3.6
import os
import sys
import random
from datetime import datetime, timedelta
from colorama import init
from termcolor import colored
import threading
import time
import vlc
from flask import Flask, url_for
import configparser
from peewee import *
from playhouse.sqlite_ext import SqliteExtDatabase

DB = SqliteExtDatabase('main.sqlite')


class BaseModel(Model):

    class Meta:
        database = DB


class Alarm(BaseModel):
    name = CharField()
    play_type = CharField()
    url = CharField()
    alarm_time = CharField()
    last_alarm_date = CharField()
    repeatDays = CharField(default="1,2,3,4,5,6,7")
    active = BooleanField(default=True)
    playing = BooleanField(default=False)


class MainMusicAlarm:
    Config = configparser.ConfigParser()
    PlayingTitle = vlc.MediaPlayer()
    ConfigFileName = 'config.ini'

    def __init__(self):
        self.Config.read(self.ConfigFileName)
        print(colored('Initializing the music Alarm checker at {}.'.format(
            datetime.now().strftime("%H:%M")), 'green'))
        self.initDB()

    def parseTime(self, time):
        h, m, s = map(int, time.split(':'))
        return datetime.now().replace(hour=h, minute=m, second=s, microsecond=0)

    def playFile(self, alarm, state):
        if state:
            self.PlayingTitle = vlc.MediaPlayer(alarm.url)
            self.PlayingTitle.play()
            alarm.last_alarm_date = datetime.now().strftime("%d.%m.%Y")
            alarm.playing = True
        else:
            self.PlayingTitle.stop()
            alarm.playing = False
        alarm.save()

    def getRandomMusicTitle(self):
        musicDirectoryFileAmount = 0
        selectedFile = ""
        selectedFileIndex = 0
        for root, dirs, files in os.walk(self.Config['ALARM']['MusicDirectory']):
            musicDirectoryFileAmount += len(files)
        selectedFileIndex = random.randint(0, musicDirectoryFileAmount - 1)
        print(colored('Get random file from the {} directory.'.format(
            self.Config['ALARM']['MusicDirectory']), 'green'))
        i = 0
        for root, dirs, files in os.walk(self.Config['ALARM']['MusicDirectory']):
            for file_ in files:
                if i == selectedFileIndex:
                    selectedFile = os.path.join(root, file_)
                i += 1
        print("Selecting {} => {}) of the {} files".format(
            selectedFile, selectedFileIndex, musicDirectoryFileAmount))
        return selectedFile

    def checkAlarms(self):
        # print(datetime.now().strftime("%H:%M"))
        for alarm in Alarm.select():
            # check if the alarm is over the current time
            if self.parseTime(alarm.alarm_time) < datetime.now():
                # check if the alarm is currently not playing
                if alarm.playing == False and alarm.active == True:
                    # check if the alarm was not already played on this day
                    if str(self.parseTime(alarm.alarm_time).weekday() + 1) in alarm.repeatDays.strip().split(','):
                        if alarm.last_alarm_date != datetime.now().strftime("%d.%m.%Y"):
                            print(colored('Triggering the alarm {} at {} for {}.'.format(
                                alarm.name, datetime.now(), self.parseTime(alarm.alarm_time)), 'green'))
                            self.manageAlarmTargets(alarm)
        time.sleep(10)
        self.checkAlarms()

    def manageAlarmTargets(self, alarm):
        if alarm.play_type == 'file':
            self.playFile(alarm, 1)
        os.system("sudo send433 10101 4 {}".format(1))

    def stopAlarm(self):
        for alarm in Alarm.select().where(Alarm.playing == 1):
            alarm.playing = False
            alarm.save()
            os.system("sudo send433 10101 4 {}".format(0))

    def initDB(self):
        DB.connect()
        Alarm.create_table(True)


app = Flask(__name__)
M = MainMusicAlarm()


@app.route('/api/v1/stopAlarm')
def api_root():
    M.stopAlarm()
    return 'Welcome'


if (__name__ == "__main__"):
    threading.Thread(target=M.checkAlarms).start()
    app.run(host=M.Config['HTTP']['ListenAddr'], port=M.Config[
            'HTTP']['ListenPort'], threaded=True)
    print(colored('The programm was terminated at {}. Stopping services...'.format(
        datetime.now()), 'red'))
    M.manageAlarm(0)
    sys.exit(os.EX_OK)
=======
#!/usr/bin/python3.6
import os
import sys
import random
from datetime import datetime, timedelta
from colorama import init
from termcolor import colored
import threading
import time
import vlc
from flask import Flask, url_for
import configparser
from peewee import *
from playhouse.sqlite_ext import SqliteExtDatabase

DB = SqliteExtDatabase('main.sqlite')


class BaseModel(Model):

    class Meta:
        database = DB


class Alarm(BaseModel):
    name = CharField()
    play_type = CharField()
    url = CharField()
    alarm_time = CharField()
    last_alarm_date = CharField()
    repeatDays = CharField(default="1,2,3,4,5,6,7")
    active = BooleanField(default=True)
    playing = BooleanField(default=False)


class MainMusicAlarm:
    Config = configparser.ConfigParser()
    PlayingTitle = vlc.MediaPlayer()
    ConfigFileName = 'config.ini'

    def __init__(self):
        self.Config.read(self.ConfigFileName)
        print(colored('Initializing the music Alarm checker at {}.'.format(
            datetime.now().strftime("%H:%M")), 'green'))
        self.initDB()

    def parseTime(self, time):
        h, m, s = map(int, time.split(':'))
        return datetime.now().replace(hour=h, minute=m, second=s, microsecond=0)

    def playFile(self, alarm, state):
        if state:
            self.PlayingTitle = vlc.MediaPlayer(alarm.url)
            self.PlayingTitle.play()
            alarm.last_alarm_date = datetime.now().strftime("%d.%m.%Y")
            alarm.playing = True
        else:
            self.PlayingTitle.stop()
            alarm.playing = False
        alarm.save()

    def getRandomMusicTitle(self):
        musicDirectoryFileAmount = 0
        selectedFile = ""
        selectedFileIndex = 0
        for root, dirs, files in os.walk(self.Config['ALARM']['MusicDirectory']):
            musicDirectoryFileAmount += len(files)
        selectedFileIndex = random.randint(0, musicDirectoryFileAmount - 1)
        print(colored('Get random file from the {} directory.'.format(
            self.Config['ALARM']['MusicDirectory']), 'green'))
        i = 0
        for root, dirs, files in os.walk(self.Config['ALARM']['MusicDirectory']):
            for file_ in files:
                if i == selectedFileIndex:
                    selectedFile = os.path.join(root, file_)
                i += 1
        print("Selecting {} => {}) of the {} files".format(
            selectedFile, selectedFileIndex, musicDirectoryFileAmount))
        return selectedFile

    def checkAlarms(self):
        # print(datetime.now().strftime("%H:%M"))
        for alarm in Alarm.select():
            # check if the alarm is over the current time
            if self.parseTime(alarm.alarm_time) < datetime.now():
                # check if the alarm is currently not playing
                if alarm.playing == False and alarm.active == True:
                    # check if the alarm was not already played on this day
                    if str(self.parseTime(alarm.alarm_time).weekday() + 1) in alarm.repeatDays.strip().split(','):
                        if alarm.last_alarm_date != datetime.now().strftime("%d.%m.%Y"):
                            print(colored('Triggering the alarm {} at {} for {}.'.format(
                                alarm.name, datetime.now(), self.parseTime(alarm.alarm_time)), 'green'))
                            self.manageAlarmTargets(alarm)
        time.sleep(10)
        self.checkAlarms()

    def manageAlarmTargets(self, alarm):
        if alarm.play_type == 'file':
            self.playFile(alarm, 1)
        os.system("sudo send433 10101 4 {}".format(1))

    def stopAlarm(self):
        for alarm in Alarm.select().where(Alarm.playing == 1):
            alarm.playing = False
            alarm.save()
            os.system("sudo send433 10101 4 {}".format(0))

    def initDB(self):
        DB.connect()
        Alarm.create_table(True)


app = Flask(__name__)
M = MainMusicAlarm()


@app.route('/api/v1/stopAlarm')
def api_root():
    M.stopAlarm()
    return 'Welcome'


if (__name__ == "__main__"):
    threading.Thread(target=M.checkAlarms).start()
    app.run(host=M.Config['HTTP']['ListenAddr'], port=M.Config[
            'HTTP']['ListenPort'], threaded=True)
    print(colored('The programm was terminated at {}. Stopping services...'.format(
        datetime.now()), 'red'))
    M.manageAlarm(0)
    sys.exit(os.EX_OK)
>>>>>>> 2d8cc64eb633f08b84b149cfc3b5e66e311cdf19