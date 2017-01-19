#!/usr/bin/python3.6
import os
import sys
import glob
import datetime
import termcolor
import threading
import logging
import subprocess
import vlc
import flask
import configparser
import gtts
import calendar
import peewee
import playhouse.sqlite_ext

DB = playhouse.sqlite_ext.SqliteExtDatabase('main.sqlite')


class BaseModel(peewee.Model):

    class Meta:
        database = DB

logging.basicConfig(
    level=logging.DEBUG,
    format='%(asctime)s:%(levelname)s:%(name)s:%(message)s',
    filename="main.log",
    filemode='a'
)


class Alarm(BaseModel):
    name = peewee.CharField()
    source = peewee.CharField()
    url = peewee.CharField()
    alarmTime = peewee.TimeField()
    created = peewee.DateTimeField(default=datetime.datetime.now)
    lastAlarm = peewee.DateTimeField(
        default=datetime.datetime.utcfromtimestamp(0))
    repeatDays = peewee.CharField()
    active = peewee.BooleanField(default=True)
    playing = peewee.BooleanField(default=False)


class MainMusicAlarm:
    Config = configparser.ConfigParser()
    VLCInstance = vlc.Instance()
    VLCPlayer = VLCInstance.media_player_new()
    ConfigFileName = 'config.ini'

    def __init__(self):
        self.Config.read(self.ConfigFileName)
        logging.info(termcolor.colored('Initializing the music Alarm checker at {}.'.format(
            datetime.datetime.now().strftime("%H:%M")), 'green'))
        self.initDB()
        # events = self.VLCPlayer.event_manager()
        # events.event_attach(
        #     vlc.EventType.MediaPlayerEndReached, self.SongFinished)

    # def SongFinished(event):
    #     global finish
    #     logging.debug("Event reports - finished")
    #     finish = 1
    #     self.stopAlarm()

    def playFile(self, alarm):
        self.VLCPlayer.set_media(self.VLCInstance.media_new(alarm.url))
        self.VLCPlayer.play()

    def playTTS(self, alarm):
        tts = gtts.gTTS(text=alarm.url, lang='de')
        tts.save("temp.mp3")
        self.VLCPlayer.set_media(self.VLCInstance.media_new("temp.mp3"))
        self.VLCPlayer.play()

    def checkAlarms(self):
        logging.warning(termcolor.colored(
            'SCHEDULER => Starting checking the alarms!', 'yellow'))
        for alarm in Alarm.select():
            # check if the current weekday is in the alarm ones and
            # check if the alarm is a single one (without repeating)
            logging.warning(termcolor.colored('1 - Passing the weekday condition?. Today is {} and the alarm will trigger => {}'.format(
                str(datetime.datetime.now().weekday() + 1), alarm.repeatDays.strip().split(',')), 'green'))
            if str(datetime.datetime.now().weekday() + 1) in alarm.repeatDays.strip().split(',') or (alarm.lastAlarm == datetime.datetime.utcfromtimestamp(0) and alarm.repeatDays == '0'):
                # check if the alarm is over the current time
                logging.warning(termcolor.colored('2 - Passing the alarm time condition?. Today is {} and the alarm will trigger => {}'.format(
                    datetime.datetime.now().time(), alarm.alarmTime), 'green'))
                if alarm.alarmTime < datetime.datetime.now().time():
                    # check if the alarm is currently not playing
                    logging.warning(termcolor.colored('3 - Passing the playing / active condition?. The alarm is active => {} and is playing => {}'.format(
                        alarm.active, alarm.playing), 'green'))
                    if alarm.playing == False and alarm.active == True:
                        # adds to the lastAlarm one day and sets the time to the normal alarm time
                        # this variable should be bigger as the
                        # datetime.datetime.now()
                        nextScheduledAlarm = alarm.lastAlarm.replace(
                            hour=alarm.alarmTime.hour, minute=alarm.alarmTime.minute, second=alarm.alarmTime.second, microsecond=0) + datetime.timedelta(days=1)
                        if ((nextScheduledAlarm < datetime.datetime.now()) or
                                (alarm.lastAlarm == datetime.datetime.utcfromtimestamp(0)) and nextScheduledAlarm > alarm.created):
                            logging.info(termcolor.colored('4 - Triggering the alarm {} at {} for {}.'.format(
                                alarm.name, datetime.datetime.now(), alarm.alarmTime), 'green'))
                            self.manageAlarmTargets(alarm)
        threading.Timer(30.0, self.checkAlarms).start()

    def manageAlarmTargets(self, alarm):
        os.system("sudo send433 10101 4 {}".format(1))
        if alarm.source == 'file' or alarm.source == 'stream':
            self.playFile(alarm)
        elif alarm.source == 'tts':
            self.playTTS(alarm)
        alarm.lastAlarm = datetime.datetime.now()
        alarm.playing = True
        alarm.save()

    def stopAlarm(self):
        for alarm in Alarm.select().where(Alarm.playing == 1):
            alarm.playing = False
            alarm.save()
        os.system("sudo send433 10101 4 {}".format(0))
        self.VLCPlayer.stop()

    def initDB(self):
        DB.connect()
        Alarm.create_table(True)


class CustomJSONEncoder(flask.json.JSONEncoder):

    def default(self, obj):
        try:
            if isinstance(obj, datetime.datetime):
                if obj.utcoffset() is not None:
                    obj = obj - obj.utcoffset()
                millis = int(
                    calendar.timegm(obj.timetuple()) * 1000 +
                    obj.microsecond / 1000
                )
                return millis
            if isinstance(obj, datetime.time):
                return obj.isoformat()
            iterable = iter(obj)
        except TypeError:
            pass
        else:
            return list(iterable)
        return flask.json.JSONEncoder.default(self, obj)

app = flask.Flask(__name__)
app.json_encoder = CustomJSONEncoder
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
                 alarmTime=datetime.datetime.strptime(
                     flask.request.json['alarmTime'], '%H:%M:%S').time(),
                 repeatDays=flask.request.json['repeatDays'],
                 active=flask.request.json['active'])
    return flask.jsonify({'success': True})


@app.route('/api/v1/deleteAlarm/<int:AlarmId>')
def deleteAlarm_page(AlarmId):
    return flask.jsonify(Alarm.get(Alarm.id == AlarmId).delete_instance())


@app.route('/api/v1/test/<int:TestId>')
def do_testing_page(TestId):
    if TestId == 1:
        M.VLCPlayer.set_media(M.VLCInstance.media_new(
            "http://listen.hardbase.fm/tunein-aacplus-pls"))
        M.VLCPlayer.play()
        os.system("sudo send433 10101 4 {}".format(1))
    if TestId == 2:
        os.system("sudo send433 10101 4 {}".format(0))
        M.VLCPlayer.stop()
    return flask.jsonify({'success': True})


@app.route('/api/v1/volume/<int:Amount>')
def get_volume_page(Amount):
    return flask.jsonify({'data': subprocess.check_output("amixer set Speaker " + str(Amount) + "% | grep % | awk '{print $5}'|sed 's/[^0-9\%]//g' | head -n 1", shell=True).decode("utf-8").rstrip('\n'), 'success': True})


@app.route('/api/v1/volume')
def set_volume_page():
    return flask.jsonify({'data': subprocess.check_output("amixer get Speaker | grep % | awk '{print $5}'|sed 's/[^0-9\%]//g' | head -n 1", shell=True).decode("utf-8").rstrip('\n'), 'success': True})


@app.route('/api/v1/modifyAlarm/<int:AlarmId>', methods=['POST'])
def modifyAlarm_page(AlarmId):
    entry = Alarm.get(Alarm.id == AlarmId)
    entry.name = flask.request.json[
        'name'] if 'name' in flask.request.json else entry.name
    entry.source = flask.request.json[
        'source'] if 'source' in flask.request.json else entry.source
    entry.url = flask.request.json[
        'url'] if 'url' in flask.request.json else entry.url
    entry.alarmTime = flask.request.json[
        'alarmTime'] if 'alarmTime' in flask.request.json else entry.alarmTime
    entry.repeatDays = flask.request.json[
        'repeatDays'] if 'repeatDays' in flask.request.json else entry.repeatDays
    entry.active = flask.request.json[
        'active'] if 'active' in flask.request.json else entry.active
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


@app.route('/settings')
def settings_page():
    return flask.render_template('settings.html', site="Settings")


class StreamToLogger(object):
    """
    Fake file-like stream object that redirects writes to a logger instance.
    """

    def __init__(self, logger, log_level=logging.DEBUG):
        self.logger = logger
        self.log_level = log_level
        self.linebuf = ''

    def write(self, buf):
        for line in buf.rstrip().splitlines():
            self.logger.log(self.log_level, line.rstrip())

    def flush(self):
        pass

if (__name__ == "__main__"):
    try:
        sys.stdout = StreamToLogger(
            logging.getLogger('STDOUT'), logging.INFO)
        sys.stderr = StreamToLogger(
            logging.getLogger('STDERR'), logging.ERROR)
        sys.setrecursionlimit(100000)
        mainThread = threading.Thread(target=M.checkAlarms)
        mainThread.daemon = True
        mainThread.start()
        app.run(host=M.Config['HTTP']['ListenAddr'], port=int(
            M.Config['HTTP']['ListenPort']), threaded=True, debug=False)
        raise Exception(
            'Flask was interrupted so the programm will terminate!')
    except Exception as e:
        logging.critical(termcolor.colored('\nThe programm was terminated at {}.\nError: {}. \nStopping services...'.format(
            datetime.datetime.now(), str(e)), 'red'))
        os.system("sudo send433 10101 4 {}".format(0))
        sys.exit(os.EX_OK)
