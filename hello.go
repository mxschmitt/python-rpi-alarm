package main
import (
	πg "grumpy"
	π_grumpyΓlibΓcalendar "grumpy/lib/calendar"
	π_grumpyΓlibΓconfigparser "grumpy/lib/configparser"
	π_grumpyΓlibΓdatetime "grumpy/lib/datetime"
	π_grumpyΓlibΓflask "grumpy/lib/flask"
	π_grumpyΓlibΓglob "grumpy/lib/glob"
	π_grumpyΓlibΓgtts "grumpy/lib/gtts"
	π_grumpyΓlibΓlogging "grumpy/lib/logging"
	π_grumpyΓlibΓos "grumpy/lib/os"
	π_grumpyΓlibΓpeewee "grumpy/lib/peewee"
	π_grumpyΓlibΓplayhouse "grumpy/lib/playhouse"
	π_grumpyΓlibΓplayhouseΓsqlite_ext "grumpy/lib/playhouse/sqlite_ext"
	π_grumpyΓlibΓsubprocess "grumpy/lib/subprocess"
	π_grumpyΓlibΓsys "grumpy/lib/sys"
	π_grumpyΓlibΓtermcolor "grumpy/lib/termcolor"
	π_grumpyΓlibΓthreading "grumpy/lib/threading"
	π_grumpyΓlibΓvlc "grumpy/lib/vlc"
	π_os "os"
)
func initModule(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
	ß := πg.InternStr("")
	ß0 := πg.InternStr("0")
	ßALARM := πg.InternStr("ALARM")
	ßAlarm := πg.InternStr("Alarm")
	ßBaseModel := πg.InternStr("BaseModel")
	ßBooleanField := πg.InternStr("BooleanField")
	ßCharField := πg.InternStr("CharField")
	ßConfig := πg.InternStr("Config")
	ßConfigFileName := πg.InternStr("ConfigFileName")
	ßConfigParser := πg.InternStr("ConfigParser")
	ßCustomJSONEncoder := πg.InternStr("CustomJSONEncoder")
	ßDB := πg.InternStr("DB")
	ßDEBUG := πg.InternStr("DEBUG")
	ßDateTimeField := πg.InternStr("DateTimeField")
	ßERROR := πg.InternStr("ERROR")
	ßEX_OK := πg.InternStr("EX_OK")
	ßEventType := πg.InternStr("EventType")
	ßException := πg.InternStr("Exception")
	ßFalse := πg.InternStr("False")
	ßFlask := πg.InternStr("Flask")
	ßHTTP := πg.InternStr("HTTP")
	ßHistory := πg.InternStr("History")
	ßINFO := πg.InternStr("INFO")
	ßInstance := πg.InternStr("Instance")
	ßIntegerField := πg.InternStr("IntegerField")
	ßJSONEncoder := πg.InternStr("JSONEncoder")
	ßListenAddr := πg.InternStr("ListenAddr")
	ßListenPort := πg.InternStr("ListenPort")
	ßM := πg.InternStr("M")
	ßMainMusicAlarm := πg.InternStr("MainMusicAlarm")
	ßMediaPlayerEndReached := πg.InternStr("MediaPlayerEndReached")
	ßMeta := πg.InternStr("Meta")
	ßModel := πg.InternStr("Model")
	ßNone := πg.InternStr("None")
	ßPOST := πg.InternStr("POST")
	ßSOCKETS := πg.InternStr("SOCKETS")
	ßSTDERR := πg.InternStr("STDERR")
	ßSTDOUT := πg.InternStr("STDOUT")
	ßSongFinished := πg.InternStr("SongFinished")
	ßSqliteExtDatabase := πg.InternStr("SqliteExtDatabase")
	ßStreamToLogger := πg.InternStr("StreamToLogger")
	ßThread := πg.InternStr("Thread")
	ßTimeField := πg.InternStr("TimeField")
	ßTimer := πg.InternStr("Timer")
	ßTrue := πg.InternStr("True")
	ßTypeError := πg.InternStr("TypeError")
	ßVLCInstance := πg.InternStr("VLCInstance")
	ßVLCPlayer := πg.InternStr("VLCPlayer")
	ß__init__ := πg.InternStr("__init__")
	ß__main__ := πg.InternStr("__main__")
	ß__metaclass__ := πg.InternStr("__metaclass__")
	ß__module__ := πg.InternStr("__module__")
	ß__name__ := πg.InternStr("__name__")
	ßa := πg.InternStr("a")
	ßactive := πg.InternStr("active")
	ßalarmTime := πg.InternStr("alarmTime")
	ßapp := πg.InternStr("app")
	ßappend := πg.InternStr("append")
	ßbasename := πg.InternStr("basename")
	ßbasicConfig := πg.InternStr("basicConfig")
	ßcalendar := πg.InternStr("calendar")
	ßcheckAlarms := πg.InternStr("checkAlarms")
	ßcheck_output := πg.InternStr("check_output")
	ßcode := πg.InternStr("code")
	ßcolored := πg.InternStr("colored")
	ßcombine := πg.InternStr("combine")
	ßconfigparser := πg.InternStr("configparser")
	ßconnect := πg.InternStr("connect")
	ßcreate := πg.InternStr("create")
	ßcreateAlarm_page := πg.InternStr("createAlarm_page")
	ßcreate_table := πg.InternStr("create_table")
	ßcreated := πg.InternStr("created")
	ßcritical := πg.InternStr("critical")
	ßdaemon := πg.InternStr("daemon")
	ßdata := πg.InternStr("data")
	ßdatabase := πg.InternStr("database")
	ßdate := πg.InternStr("date")
	ßdatetime := πg.InternStr("datetime")
	ßde := πg.InternStr("de")
	ßdebug := πg.InternStr("debug")
	ßdecode := πg.InternStr("decode")
	ßdefault := πg.InternStr("default")
	ßdefault_serve := πg.InternStr("default_serve")
	ßdeleteAlarm_page := πg.InternStr("deleteAlarm_page")
	ßdelete_instance := πg.InternStr("delete_instance")
	ßdicts := πg.InternStr("dicts")
	ßdo_testing_page := πg.InternStr("do_testing_page")
	ßduration := πg.InternStr("duration")
	ße := πg.InternStr("e")
	ßevent_attach := πg.InternStr("event_attach")
	ßevent_manager := πg.InternStr("event_manager")
	ßexit := πg.InternStr("exit")
	ßfile := πg.InternStr("file")
	ßfilename := πg.InternStr("filename")
	ßflask := πg.InternStr("flask")
	ßflush := πg.InternStr("flush")
	ßformat := πg.InternStr("format")
	ßgTTS := πg.InternStr("gTTS")
	ßget := πg.InternStr("get")
	ßgetAlarms := πg.InternStr("getAlarms")
	ßgetFiles_page := πg.InternStr("getFiles_page")
	ßgetHistory := πg.InternStr("getHistory")
	ßgetLogger := πg.InternStr("getLogger")
	ßget_volume_page := πg.InternStr("get_volume_page")
	ßglob := πg.InternStr("glob")
	ßgreen := πg.InternStr("green")
	ßgtts := πg.InternStr("gtts")
	ßhour := πg.InternStr("hour")
	ßid := πg.InternStr("id")
	ßiglob := πg.InternStr("iglob")
	ßinfo := πg.InternStr("info")
	ßinitDB := πg.InternStr("initDB")
	ßint := πg.InternStr("int")
	ßisinstance := πg.InternStr("isinstance")
	ßisoformat := πg.InternStr("isoformat")
	ßiter := πg.InternStr("iter")
	ßjson := πg.InternStr("json")
	ßjson_encoder := πg.InternStr("json_encoder")
	ßjsonify := πg.InternStr("jsonify")
	ßlastAlarm := πg.InternStr("lastAlarm")
	ßlinebuf := πg.InternStr("linebuf")
	ßlist := πg.InternStr("list")
	ßlog := πg.InternStr("log")
	ßlog_level := πg.InternStr("log_level")
	ßlogger := πg.InternStr("logger")
	ßlogging := πg.InternStr("logging")
	ßmainThread := πg.InternStr("mainThread")
	ßmain_page := πg.InternStr("main_page")
	ßmanageAlarmTargets := πg.InternStr("manageAlarmTargets")
	ßmedia_new := πg.InternStr("media_new")
	ßmedia_player_new := πg.InternStr("media_player_new")
	ßmicrosecond := πg.InternStr("microsecond")
	ßminute := πg.InternStr("minute")
	ßmodifyAlarm_page := πg.InternStr("modifyAlarm_page")
	ßmusicdirectory := πg.InternStr("musicdirectory")
	ßname := πg.InternStr("name")
	ßnow := πg.InternStr("now")
	ßobject := πg.InternStr("object")
	ßos := πg.InternStr("os")
	ßpath := πg.InternStr("path")
	ßpeewee := πg.InternStr("peewee")
	ßplay := πg.InternStr("play")
	ßplayFile := πg.InternStr("playFile")
	ßplayTTS := πg.InternStr("playTTS")
	ßplayhouse := πg.InternStr("playhouse")
	ßplaying := πg.InternStr("playing")
	ßpowerSocket := πg.InternStr("powerSocket")
	ßread := πg.InternStr("read")
	ßred := πg.InternStr("red")
	ßrepeatDays := πg.InternStr("repeatDays")
	ßreplace := πg.InternStr("replace")
	ßrequest := πg.InternStr("request")
	ßroute := πg.InternStr("route")
	ßrstrip := πg.InternStr("rstrip")
	ßrun := πg.InternStr("run")
	ßsave := πg.InternStr("save")
	ßsecond := πg.InternStr("second")
	ßselect := πg.InternStr("select")
	ßsend_from_directory := πg.InternStr("send_from_directory")
	ßset_media := πg.InternStr("set_media")
	ßset_volume_page := πg.InternStr("set_volume_page")
	ßsetrecursionlimit := πg.InternStr("setrecursionlimit")
	ßsource := πg.InternStr("source")
	ßsplit := πg.InternStr("split")
	ßsplitlines := πg.InternStr("splitlines")
	ßsqlite_ext := πg.InternStr("sqlite_ext")
	ßstart := πg.InternStr("start")
	ßstatic := πg.InternStr("static")
	ßstderr := πg.InternStr("stderr")
	ßstdout := πg.InternStr("stdout")
	ßstop := πg.InternStr("stop")
	ßstopAlarm := πg.InternStr("stopAlarm")
	ßstopTime := πg.InternStr("stopTime")
	ßstr := πg.InternStr("str")
	ßstream := πg.InternStr("stream")
	ßstrftime := πg.InternStr("strftime")
	ßstrip := πg.InternStr("strip")
	ßstrptime := πg.InternStr("strptime")
	ßsubprocess := πg.InternStr("subprocess")
	ßsuccess := πg.InternStr("success")
	ßsys := πg.InternStr("sys")
	ßsystem := πg.InternStr("system")
	ßtermcolor := πg.InternStr("termcolor")
	ßthreading := πg.InternStr("threading")
	ßtime := πg.InternStr("time")
	ßtimedelta := πg.InternStr("timedelta")
	ßtimegm := πg.InternStr("timegm")
	ßtimetuple := πg.InternStr("timetuple")
	ßtoday := πg.InternStr("today")
	ßtotal_seconds := πg.InternStr("total_seconds")
	ßtriggerTime := πg.InternStr("triggerTime")
	ßtts := πg.InternStr("tts")
	ßurl := πg.InternStr("url")
	ßutcfromtimestamp := πg.InternStr("utcfromtimestamp")
	ßutcoffset := πg.InternStr("utcoffset")
	ßvlc := πg.InternStr("vlc")
	ßwarning := πg.InternStr("warning")
	ßweekday := πg.InternStr("weekday")
	ßwhere := πg.InternStr("where")
	ßwrite := πg.InternStr("write")
	ßyellow := πg.InternStr("yellow")
	var πTemp001 *πg.Object
	_ = πTemp001
	var πTemp002 []*πg.Object
	_ = πTemp002
	var πTemp003 *πg.Object
	_ = πTemp003
	var πTemp004 *πg.Dict
	_ = πTemp004
	var πTemp005 *πg.Object
	_ = πTemp005
	var πTemp006 *πg.Object
	_ = πTemp006
	var πTemp007 πg.KWArgs
	_ = πTemp007
	var πTemp008 []πg.Param
	_ = πTemp008
	var πTemp009 []*πg.Object
	_ = πTemp009
	var πTemp010 []*πg.Object
	_ = πTemp010
	var πTemp011 *πg.Object
	_ = πTemp011
	var πTemp012 *πg.Object
	_ = πTemp012
	var πTemp013 *πg.Object
	_ = πTemp013
	var πTemp014 *πg.Object
	_ = πTemp014
	var πTemp015 *πg.Object
	_ = πTemp015
	var πTemp016 *πg.Object
	_ = πTemp016
	var πTemp017 *πg.Object
	_ = πTemp017
	var πTemp018 *πg.Object
	_ = πTemp018
	var πTemp019 *πg.Object
	_ = πTemp019
	var πTemp020 *πg.Object
	_ = πTemp020
	var πTemp021 *πg.Object
	_ = πTemp021
	var πTemp022 bool
	_ = πTemp022
	var πTemp023 *πg.Object
	_ = πTemp023
	var πTemp024 *πg.Object
	_ = πTemp024
	var πTemp025 *πg.Object
	_ = πTemp025
	var πTemp026 *πg.Object
	_ = πTemp026
	var πTemp027 *πg.BaseException
	_ = πTemp027
	var πTemp028 *πg.Traceback
	_ = πTemp028
	var πTemp029 []*πg.Object
	_ = πTemp029
	var πE *πg.BaseException; _ = πE
	for ; πF.State() >= 0; πF.PopCheckpoint() {
		switch πF.State() {
		case 0:
		case 4: goto Label4
		default: panic("unexpected function state")
		}
		// line 2: import os
		πF.SetLineno(2)
		if πTemp002, πE = πg.ImportModule(πF, "os", []*πg.Code{π_grumpyΓlibΓos.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßos.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 3: import sys
		πF.SetLineno(3)
		if πTemp002, πE = πg.ImportModule(πF, "sys", []*πg.Code{π_grumpyΓlibΓsys.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 4: import glob
		πF.SetLineno(4)
		if πTemp002, πE = πg.ImportModule(πF, "glob", []*πg.Code{π_grumpyΓlibΓglob.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßglob.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 5: import datetime
		πF.SetLineno(5)
		if πTemp002, πE = πg.ImportModule(πF, "datetime", []*πg.Code{π_grumpyΓlibΓdatetime.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßdatetime.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 6: import termcolor
		πF.SetLineno(6)
		if πTemp002, πE = πg.ImportModule(πF, "termcolor", []*πg.Code{π_grumpyΓlibΓtermcolor.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßtermcolor.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 7: import threading
		πF.SetLineno(7)
		if πTemp002, πE = πg.ImportModule(πF, "threading", []*πg.Code{π_grumpyΓlibΓthreading.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßthreading.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 8: import logging
		πF.SetLineno(8)
		if πTemp002, πE = πg.ImportModule(πF, "logging", []*πg.Code{π_grumpyΓlibΓlogging.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßlogging.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 9: import subprocess
		πF.SetLineno(9)
		if πTemp002, πE = πg.ImportModule(πF, "subprocess", []*πg.Code{π_grumpyΓlibΓsubprocess.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßsubprocess.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 10: import vlc
		πF.SetLineno(10)
		if πTemp002, πE = πg.ImportModule(πF, "vlc", []*πg.Code{π_grumpyΓlibΓvlc.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßvlc.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 11: import flask
		πF.SetLineno(11)
		if πTemp002, πE = πg.ImportModule(πF, "flask", []*πg.Code{π_grumpyΓlibΓflask.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßflask.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 12: import configparser
		πF.SetLineno(12)
		if πTemp002, πE = πg.ImportModule(πF, "configparser", []*πg.Code{π_grumpyΓlibΓconfigparser.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßconfigparser.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 13: import gtts
		πF.SetLineno(13)
		if πTemp002, πE = πg.ImportModule(πF, "gtts", []*πg.Code{π_grumpyΓlibΓgtts.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßgtts.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 14: import calendar
		πF.SetLineno(14)
		if πTemp002, πE = πg.ImportModule(πF, "calendar", []*πg.Code{π_grumpyΓlibΓcalendar.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßcalendar.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 15: import peewee
		πF.SetLineno(15)
		if πTemp002, πE = πg.ImportModule(πF, "peewee", []*πg.Code{π_grumpyΓlibΓpeewee.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßpeewee.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 16: import playhouse.sqlite_ext
		πF.SetLineno(16)
		if πTemp002, πE = πg.ImportModule(πF, "playhouse.sqlite_ext", []*πg.Code{π_grumpyΓlibΓplayhouse.Code, π_grumpyΓlibΓplayhouseΓsqlite_ext.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßplayhouse.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 18: DB = playhouse.sqlite_ext.SqliteExtDatabase('main.sqlite')
		πF.SetLineno(18)
		πTemp002 = πF.MakeArgs(1)
		πTemp002[0] = πg.NewStr("main.sqlite").ToObject()
		if πTemp001, πE = πg.ResolveGlobal(πF, ßplayhouse); πE != nil {
			continue
		}
		if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßsqlite_ext, nil); πE != nil {
			continue
		}
		if πTemp001, πE = πg.GetAttr(πF, πTemp003, ßSqliteExtDatabase, nil); πE != nil {
			continue
		}
		if πTemp003, πE = πTemp001.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßDB.ToObject(), πTemp003); πE != nil {
			continue
		}
		// line 21: class BaseModel(peewee.Model):
		πF.SetLineno(21)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp005, πE = πg.ResolveGlobal(πF, ßpeewee); πE != nil {
			continue
		}
		if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßModel, nil); πE != nil {
			continue
		}
		πTemp002[0] = πTemp006
		πTemp004 = πg.NewDict()
		if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
			continue
		}
		_, πE = πg.NewCode("BaseModel", "main.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp004
			_ = πClass
			var πTemp001 *πg.Dict
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 []*πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 23: class Meta:
				πF.SetLineno(23)
				πTemp003 = make([]*πg.Object, 0)
				πTemp001 = πg.NewDict()
				if πTemp002, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
					continue
				}
				if πE = πTemp001.SetItem(πF, ß__module__.ToObject(), πTemp002); πE != nil {
					continue
				}
				_, πE = πg.NewCode("Meta", "main.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
					πClass := πTemp001
					_ = πClass
					var πTemp001 *πg.Object
					_ = πTemp001
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 24: database = DB
						πF.SetLineno(24)
						if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßDB); πE != nil {
							continue
						}
						if πE = πClass.SetItem(πF, ßdatabase.ToObject(), πTemp001); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}).Eval(πF, πF.Globals(), nil, nil)
				if πE != nil {
					continue
				}
				if πTemp004, πE = πTemp001.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
					continue
				}
				if πTemp004 == nil {
					πTemp004 = πg.TypeType.ToObject()
				}
				if πTemp005, πE = πTemp004.Call(πF, []*πg.Object{πg.NewStr("Meta").ToObject(), πg.NewTuple(πTemp003...).ToObject(), πTemp001.ToObject()}, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßMeta.ToObject(), πTemp005); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}).Eval(πF, πF.Globals(), nil, nil)
		if πE != nil {
			continue
		}
		if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			continue
		}
		if πTemp003 == nil {
			πTemp003 = πg.TypeType.ToObject()
		}
		if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("BaseModel").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßBaseModel.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 26: logging.basicConfig(
		πF.SetLineno(26)
		if πTemp001, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
			continue
		}
		if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßDEBUG, nil); πE != nil {
			continue
		}
		πTemp007 = πg.KWArgs{
			{"level", πTemp003},
			{"format", πg.NewStr("%(asctime)s:%(levelname)s:%(name)s:%(message)s").ToObject()},
			{"filename", πg.NewStr("main.log").ToObject()},
			{"filemode", ßa.ToObject()},
		}
		if πTemp001, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
			continue
		}
		if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßbasicConfig, nil); πE != nil {
			continue
		}
		if πTemp001, πE = πTemp003.Call(πF, nil, πTemp007); πE != nil {
			continue
		}
		// line 34: class Alarm(BaseModel):
		πF.SetLineno(34)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseModel); πE != nil {
			continue
		}
		πTemp002[0] = πTemp005
		πTemp004 = πg.NewDict()
		if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
			continue
		}
		_, πE = πg.NewCode("Alarm", "main.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp004
			_ = πClass
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 πg.KWArgs
			_ = πTemp003
			var πTemp004 []*πg.Object
			_ = πTemp004
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 35: name = peewee.CharField()
				πF.SetLineno(35)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßCharField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßname.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 36: source = peewee.CharField()
				πF.SetLineno(36)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßCharField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßsource.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 37: url = peewee.CharField()
				πF.SetLineno(37)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßCharField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßurl.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 38: alarmTime = peewee.TimeField()
				πF.SetLineno(38)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßTimeField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßalarmTime.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 39: created = peewee.DateTimeField(default=datetime.datetime.now)
				πF.SetLineno(39)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdatetime); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdatetime, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßnow, nil); πE != nil {
					continue
				}
				πTemp003 = πg.KWArgs{
					{"default", πTemp001},
				}
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßDateTimeField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, πTemp003); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßcreated.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 40: lastAlarm = peewee.DateTimeField(
				πF.SetLineno(40)
				πTemp004 = πF.MakeArgs(1)
				πTemp004[0] = πg.NewInt(0).ToObject()
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßdatetime); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdatetime, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßutcfromtimestamp, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003 = πg.KWArgs{
					{"default", πTemp002},
				}
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßDateTimeField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, πTemp003); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßlastAlarm.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 42: repeatDays = peewee.CharField()
				πF.SetLineno(42)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßCharField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßrepeatDays.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 43: active = peewee.BooleanField(default=True)
				πF.SetLineno(43)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßTrue); πE != nil {
					continue
				}
				πTemp003 = πg.KWArgs{
					{"default", πTemp001},
				}
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßBooleanField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, πTemp003); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßactive.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 44: playing = peewee.BooleanField(default=False)
				πF.SetLineno(44)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßFalse); πE != nil {
					continue
				}
				πTemp003 = πg.KWArgs{
					{"default", πTemp001},
				}
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßBooleanField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, πTemp003); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßplaying.ToObject(), πTemp001); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}).Eval(πF, πF.Globals(), nil, nil)
		if πE != nil {
			continue
		}
		if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			continue
		}
		if πTemp003 == nil {
			πTemp003 = πg.TypeType.ToObject()
		}
		if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("Alarm").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßAlarm.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 47: class History(BaseModel):
		πF.SetLineno(47)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp005, πE = πg.ResolveGlobal(πF, ßBaseModel); πE != nil {
			continue
		}
		πTemp002[0] = πTemp005
		πTemp004 = πg.NewDict()
		if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
			continue
		}
		_, πE = πg.NewCode("History", "main.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp004
			_ = πClass
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 48: triggerTime = peewee.TimeField()
				πF.SetLineno(48)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßTimeField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßtriggerTime.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 49: stopTime = peewee.TimeField()
				πF.SetLineno(49)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßTimeField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßstopTime.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 50: duration = peewee.IntegerField()
				πF.SetLineno(50)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßpeewee); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßIntegerField, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßduration.ToObject(), πTemp001); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}).Eval(πF, πF.Globals(), nil, nil)
		if πE != nil {
			continue
		}
		if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			continue
		}
		if πTemp003 == nil {
			πTemp003 = πg.TypeType.ToObject()
		}
		if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("History").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßHistory.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 53: class MainMusicAlarm:
		πF.SetLineno(53)
		πTemp002 = make([]*πg.Object, 0)
		πTemp004 = πg.NewDict()
		if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
			continue
		}
		_, πE = πg.NewCode("MainMusicAlarm", "main.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp004
			_ = πClass
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 []πg.Param
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πTemp006 *πg.Object
			_ = πTemp006
			var πTemp007 *πg.Object
			_ = πTemp007
			var πTemp008 *πg.Object
			_ = πTemp008
			var πTemp009 *πg.Object
			_ = πTemp009
			var πTemp010 *πg.Object
			_ = πTemp010
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 54: Config = configparser.ConfigParser()
				πF.SetLineno(54)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßconfigparser); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßConfigParser, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßConfig.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 55: VLCInstance = vlc.Instance()
				πF.SetLineno(55)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßvlc); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßInstance, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßVLCInstance.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 56: VLCPlayer = VLCInstance.media_player_new()
				πF.SetLineno(56)
				if πTemp001, πE = πg.ResolveClass(πF, πClass, nil, ßVLCInstance); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmedia_player_new, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πE = πClass.SetItem(πF, ßVLCPlayer.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 57: ConfigFileName = 'config.ini'
				πF.SetLineno(57)
				if πE = πClass.SetItem(πF, ßConfigFileName.ToObject(), πg.NewStr("config.ini").ToObject()); πE != nil {
					continue
				}
				// line 59: def __init__(self):
				πF.SetLineno(59)
				πTemp003 = make([]πg.Param, 1)
				πTemp003[0] = πg.Param{Name: "self", Def: nil}
				πTemp001 = πg.NewFunction(πg.NewCode("__init__", "main.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µevents *πg.Object = πg.UnboundLocal; _ = µevents
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 []*πg.Object
					_ = πTemp004
					var πTemp005 []*πg.Object
					_ = πTemp005
					var πTemp006 []*πg.Object
					_ = πTemp006
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 60: self.Config.read(self.ConfigFileName)
						πF.SetLineno(60)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßConfigFileName, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßConfig, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßread, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 61: logging.info(termcolor.colored('Initializing the music Alarm checker at {}.'.format(
						πF.SetLineno(61)
						πTemp001 = πF.MakeArgs(1)
						πTemp004 = πF.MakeArgs(2)
						πTemp005 = πF.MakeArgs(1)
						πTemp006 = πF.MakeArgs(1)
						πTemp006[0] = πg.NewStr("%H:%M").ToObject()
						if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßnow, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßstrftime, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						πTemp005[0] = πTemp003
						if πTemp002, πE = πg.GetAttr(πF, πg.NewStr("Initializing the music Alarm checker at {}.").ToObject(), ßformat, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						πTemp004[0] = πTemp003
						πTemp004[1] = ßgreen.ToObject()
						if πTemp002, πE = πg.ResolveGlobal(πF, ßtermcolor); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßcolored, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						πTemp001[0] = πTemp002
						if πTemp002, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßinfo, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 63: self.initDB()
						πF.SetLineno(63)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßinitDB, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						// line 64: events = self.VLCPlayer.event_manager()
						πF.SetLineno(64)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßVLCPlayer, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßevent_manager, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						µevents = πTemp002
						// line 65: events.event_attach(
						πF.SetLineno(65)
						πTemp001 = πF.MakeArgs(2)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßvlc); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßEventType, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßMediaPlayerEndReached, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßSongFinished, nil); πE != nil {
							continue
						}
						πTemp001[1] = πTemp002
						if πE = πg.CheckLocal(πF, µevents, "events"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µevents, ßevent_attach, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 68: def SongFinished(self, event):
				πF.SetLineno(68)
				πTemp003 = make([]πg.Param, 2)
				πTemp003[0] = πg.Param{Name: "self", Def: nil}
				πTemp003[1] = πg.Param{Name: "event", Def: nil}
				πTemp002 = πg.NewFunction(πg.NewCode("SongFinished", "main.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µevent *πg.Object = πArgs[1]; _ = µevent
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 69: logging.debug("Powering off the socket because the track ends...")
						πF.SetLineno(69)
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = πg.NewStr("Powering off the socket because the track ends...").ToObject()
						if πTemp002, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdebug, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 70: self.powerSocket(True)
						πF.SetLineno(70)
						πTemp001 = πF.MakeArgs(1)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßpowerSocket, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 71: self.stopAlarm()
						πF.SetLineno(71)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßstopAlarm, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßSongFinished.ToObject(), πTemp002); πE != nil {
					continue
				}
				// line 73: def playFile(self, alarm):
				πF.SetLineno(73)
				πTemp003 = make([]πg.Param, 2)
				πTemp003[0] = πg.Param{Name: "self", Def: nil}
				πTemp003[1] = πg.Param{Name: "alarm", Def: nil}
				πTemp004 = πg.NewFunction(πg.NewCode("playFile", "main.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µalarm *πg.Object = πArgs[1]; _ = µalarm
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 []*πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 *πg.Object
					_ = πTemp004
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 74: self.VLCPlayer.set_media(self.VLCInstance.media_new(alarm.url))
						πF.SetLineno(74)
						πTemp001 = πF.MakeArgs(1)
						πTemp002 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µalarm, ßurl, nil); πE != nil {
							continue
						}
						πTemp002[0] = πTemp003
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µself, ßVLCInstance, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßmedia_new, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πTemp001[0] = πTemp003
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µself, ßVLCPlayer, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßset_media, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 75: self.VLCPlayer.play()
						πF.SetLineno(75)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µself, ßVLCPlayer, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßplay, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßplayFile.ToObject(), πTemp004); πE != nil {
					continue
				}
				// line 77: def playTTS(self, alarm):
				πF.SetLineno(77)
				πTemp003 = make([]πg.Param, 2)
				πTemp003[0] = πg.Param{Name: "self", Def: nil}
				πTemp003[1] = πg.Param{Name: "alarm", Def: nil}
				πTemp005 = πg.NewFunction(πg.NewCode("playTTS", "main.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µalarm *πg.Object = πArgs[1]; _ = µalarm
					var µtts *πg.Object = πg.UnboundLocal; _ = µtts
					var πTemp001 *πg.Object
					_ = πTemp001
					var πTemp002 πg.KWArgs
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 []*πg.Object
					_ = πTemp004
					var πTemp005 []*πg.Object
					_ = πTemp005
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 78: tts = gtts.gTTS(text=alarm.url, lang='de')
						πF.SetLineno(78)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, µalarm, ßurl, nil); πE != nil {
							continue
						}
						πTemp002 = πg.KWArgs{
							{"text", πTemp001},
							{"lang", ßde.ToObject()},
						}
						if πTemp001, πE = πg.ResolveGlobal(πF, ßgtts); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßgTTS, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp003.Call(πF, nil, πTemp002); πE != nil {
							continue
						}
						µtts = πTemp001
						// line 79: tts.save("temp.mp3")
						πF.SetLineno(79)
						πTemp004 = πF.MakeArgs(1)
						πTemp004[0] = πg.NewStr("temp.mp3").ToObject()
						if πE = πg.CheckLocal(πF, µtts, "tts"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, µtts, ßsave, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						// line 80: self.VLCPlayer.set_media(self.VLCInstance.media_new("temp.mp3"))
						πF.SetLineno(80)
						πTemp004 = πF.MakeArgs(1)
						πTemp005 = πF.MakeArgs(1)
						πTemp005[0] = πg.NewStr("temp.mp3").ToObject()
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, µself, ßVLCInstance, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßmedia_new, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp003.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						πTemp004[0] = πTemp001
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, µself, ßVLCPlayer, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßset_media, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp003.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						// line 81: self.VLCPlayer.play()
						πF.SetLineno(81)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, µself, ßVLCPlayer, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßplay, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßplayTTS.ToObject(), πTemp005); πE != nil {
					continue
				}
				// line 83: def checkAlarms(self):
				πF.SetLineno(83)
				πTemp003 = make([]πg.Param, 1)
				πTemp003[0] = πg.Param{Name: "self", Def: nil}
				πTemp006 = πg.NewFunction(πg.NewCode("checkAlarms", "main.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µalarm *πg.Object = πg.UnboundLocal; _ = µalarm
					var µnextScheduledAlarm *πg.Object = πg.UnboundLocal; _ = µnextScheduledAlarm
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 []*πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 *πg.Object
					_ = πTemp004
					var πTemp005 *πg.Object
					_ = πTemp005
					var πTemp006 []*πg.Object
					_ = πTemp006
					var πTemp007 []*πg.Object
					_ = πTemp007
					var πTemp008 *πg.Object
					_ = πTemp008
					var πTemp009 *πg.Object
					_ = πTemp009
					var πTemp010 *πg.Object
					_ = πTemp010
					var πTemp011 bool
					_ = πTemp011
					var πTemp012 *πg.Object
					_ = πTemp012
					var πTemp013 *πg.Object
					_ = πTemp013
					var πTemp014 bool
					_ = πTemp014
					var πTemp015 *πg.Object
					_ = πTemp015
					var πTemp016 πg.KWArgs
					_ = πTemp016
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 84: logging.warning(termcolor.colored(
						πF.SetLineno(84)
						πTemp001 = πF.MakeArgs(1)
						πTemp002 = πF.MakeArgs(2)
						πTemp002[0] = πg.NewStr("SCHEDULER => Starting checking the alarms!").ToObject()
						πTemp002[1] = ßyellow.ToObject()
						if πTemp003, πE = πg.ResolveGlobal(πF, ßtermcolor); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßcolored, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πTemp001[0] = πTemp003
						if πTemp003, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßwarning, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 86: for alarm in Alarm.select():
						πF.SetLineno(86)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßAlarm); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßselect, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp004, πE = πg.Iter(πF, πTemp003); πE != nil {
							continue
						}
					Label1:
						if πTemp005, πE = πg.Next(πF, πTemp004); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
								continue
							}
							if !isStop {
								continue
							}
							πE = nil
							πF.RestoreExc(nil, nil)
							goto Label2
						}
						µalarm = πTemp005
						// line 89: logging.warning(termcolor.colored('1 - Passing the weekday condition?. Today is {} and the alarm will trigger => {}'.format(
						πF.SetLineno(89)
						πTemp001 = πF.MakeArgs(1)
						πTemp002 = πF.MakeArgs(2)
						πTemp006 = πF.MakeArgs(2)
						πTemp007 = πF.MakeArgs(1)
						if πTemp009, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp010, ßnow, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp010, ßweekday, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.Add(πF, πTemp010, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						πTemp007[0] = πTemp008
						if πTemp008, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, πTemp007, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp007)
						πTemp006[0] = πTemp009
						πTemp007 = πF.MakeArgs(1)
						πTemp007[0] = πg.NewStr(",").ToObject()
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, µalarm, ßrepeatDays, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßstrip, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßsplit, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, πTemp007, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp007)
						πTemp006[1] = πTemp008
						if πTemp008, πE = πg.GetAttr(πF, πg.NewStr("1 - Passing the weekday condition?. Today is {} and the alarm will trigger => {}").ToObject(), ßformat, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						πTemp002[0] = πTemp009
						πTemp002[1] = ßgreen.ToObject()
						if πTemp008, πE = πg.ResolveGlobal(πF, ßtermcolor); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßcolored, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πTemp001[0] = πTemp008
						if πTemp008, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßwarning, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πTemp001 = πF.MakeArgs(1)
						if πTemp012, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetAttr(πF, πTemp013, ßnow, nil); πE != nil {
							continue
						}
						if πTemp013, πE = πTemp012.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetAttr(πF, πTemp013, ßweekday, nil); πE != nil {
							continue
						}
						if πTemp013, πE = πTemp012.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Add(πF, πTemp013, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						πTemp001[0] = πTemp010
						if πTemp010, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
							continue
						}
						if πTemp012, πE = πTemp010.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = πg.NewStr(",").ToObject()
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, µalarm, ßrepeatDays, nil); πE != nil {
							continue
						}
						if πTemp013, πE = πg.GetAttr(πF, πTemp010, ßstrip, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πTemp013.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp013, πE = πg.GetAttr(πF, πTemp010, ßsplit, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πTemp013.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp014, πE = πg.Contains(πF, πTemp010, πTemp012); πE != nil {
							continue
						}
						πTemp009 = πg.GetBool(πTemp014).ToObject()
						πTemp008 = πTemp009
						if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
							continue
						}
						if πTemp011 {
							goto Label3
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetAttr(πF, µalarm, ßlastAlarm, nil); πE != nil {
							continue
						}
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = πg.NewInt(0).ToObject()
						if πTemp013, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp015, πE = πg.GetAttr(πF, πTemp013, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp013, πE = πg.GetAttr(πF, πTemp015, ßutcfromtimestamp, nil); πE != nil {
							continue
						}
						if πTemp015, πE = πTemp013.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp010, πE = πg.Eq(πF, πTemp012, πTemp015); πE != nil {
							continue
						}
						πTemp009 = πTemp010
						if πTemp014, πE = πg.IsTrue(πF, πTemp009); πE != nil {
							continue
						}
						if !πTemp014 {
							goto Label4
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetAttr(πF, µalarm, ßrepeatDays, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πg.Eq(πF, πTemp012, ß0.ToObject()); πE != nil {
							continue
						}
						πTemp009 = πTemp010
					Label4:
						πTemp008 = πTemp009
					Label3:
						if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
							continue
						}
						if πTemp011 {
							goto Label5
						}
						goto Label6
						// line 91: if str(datetime.datetime.now().weekday() + 1) in alarm.repeatDays.strip().split(',') or (alarm.lastAlarm == datetime.datetime.utcfromtimestamp(0) and alarm.repeatDays == '0'):
						πF.SetLineno(91)
					Label5:
						// line 93: logging.warning(termcolor.colored('2 - Passing the alarm time condition?. Today is {} and the alarm will trigger => {}'.format(
						πF.SetLineno(93)
						πTemp001 = πF.MakeArgs(1)
						πTemp002 = πF.MakeArgs(2)
						πTemp006 = πF.MakeArgs(2)
						if πTemp008, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp009, ßnow, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp009, ßtime, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp006[0] = πTemp009
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, µalarm, ßalarmTime, nil); πE != nil {
							continue
						}
						πTemp006[1] = πTemp008
						if πTemp008, πE = πg.GetAttr(πF, πg.NewStr("2 - Passing the alarm time condition?. Today is {} and the alarm will trigger => {}").ToObject(), ßformat, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						πTemp002[0] = πTemp009
						πTemp002[1] = ßgreen.ToObject()
						if πTemp008, πE = πg.ResolveGlobal(πF, ßtermcolor); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßcolored, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πTemp001[0] = πTemp008
						if πTemp008, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßwarning, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, µalarm, ßalarmTime, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetAttr(πF, πTemp010, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, πTemp012, ßnow, nil); πE != nil {
							continue
						}
						if πTemp012, πE = πTemp010.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, πTemp012, ßtime, nil); πE != nil {
							continue
						}
						if πTemp012, πE = πTemp010.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.LT(πF, πTemp009, πTemp012); πE != nil {
							continue
						}
						if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
							continue
						}
						if πTemp011 {
							goto Label7
						}
						goto Label8
						// line 95: if alarm.alarmTime < datetime.datetime.now().time():
						πF.SetLineno(95)
					Label7:
						// line 97: logging.warning(termcolor.colored('3 - Passing the playing / active condition?. The alarm is active => {} and is playing => {}'.format(
						πF.SetLineno(97)
						πTemp001 = πF.MakeArgs(1)
						πTemp002 = πF.MakeArgs(2)
						πTemp006 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, µalarm, ßactive, nil); πE != nil {
							continue
						}
						πTemp006[0] = πTemp008
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, µalarm, ßplaying, nil); πE != nil {
							continue
						}
						πTemp006[1] = πTemp008
						if πTemp008, πE = πg.GetAttr(πF, πg.NewStr("3 - Passing the playing / active condition?. The alarm is active => {} and is playing => {}").ToObject(), ßformat, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						πTemp002[0] = πTemp009
						πTemp002[1] = ßgreen.ToObject()
						if πTemp008, πE = πg.ResolveGlobal(πF, ßtermcolor); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßcolored, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πTemp001[0] = πTemp008
						if πTemp008, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßwarning, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, µalarm, ßplaying, nil); πE != nil {
							continue
						}
						if πTemp012, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						if πTemp009, πE = πg.Eq(πF, πTemp010, πTemp012); πE != nil {
							continue
						}
						πTemp008 = πTemp009
						if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
							continue
						}
						if !πTemp011 {
							goto Label9
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, µalarm, ßactive, nil); πE != nil {
							continue
						}
						if πTemp012, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						if πTemp009, πE = πg.Eq(πF, πTemp010, πTemp012); πE != nil {
							continue
						}
						πTemp008 = πTemp009
					Label9:
						if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
							continue
						}
						if πTemp011 {
							goto Label10
						}
						goto Label11
						// line 99: if alarm.playing == False and alarm.active == True:
						πF.SetLineno(99)
					Label10:
						// line 103: nextScheduledAlarm = alarm.lastAlarm.replace(
						πF.SetLineno(103)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, µalarm, ßalarmTime, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßhour, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, µalarm, ßalarmTime, nil); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetAttr(πF, πTemp009, ßminute, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, µalarm, ßalarmTime, nil); πE != nil {
							continue
						}
						if πTemp013, πE = πg.GetAttr(πF, πTemp009, ßsecond, nil); πE != nil {
							continue
						}
						πTemp016 = πg.KWArgs{
							{"hour", πTemp010},
							{"minute", πTemp012},
							{"second", πTemp013},
							{"microsecond", πg.NewInt(0).ToObject()},
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, µalarm, ßlastAlarm, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßreplace, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp010.Call(πF, nil, πTemp016); πE != nil {
							continue
						}
						πTemp016 = πg.KWArgs{
							{"days", πg.NewInt(1).ToObject()},
						}
						if πTemp010, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetAttr(πF, πTemp010, ßtimedelta, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πTemp012.Call(πF, nil, πTemp016); πE != nil {
							continue
						}
						if πTemp008, πE = πg.Add(πF, πTemp009, πTemp010); πE != nil {
							continue
						}
						µnextScheduledAlarm = πTemp008
						if πE = πg.CheckLocal(πF, µnextScheduledAlarm, "nextScheduledAlarm"); πE != nil {
							continue
						}
						if πTemp010, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetAttr(πF, πTemp010, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, πTemp012, ßnow, nil); πE != nil {
							continue
						}
						if πTemp012, πE = πTemp010.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πg.LT(πF, µnextScheduledAlarm, πTemp012); πE != nil {
							continue
						}
						πTemp008 = πTemp009
						if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
							continue
						}
						if πTemp011 {
							goto Label12
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetAttr(πF, µalarm, ßlastAlarm, nil); πE != nil {
							continue
						}
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = πg.NewInt(0).ToObject()
						if πTemp013, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp015, πE = πg.GetAttr(πF, πTemp013, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp013, πE = πg.GetAttr(πF, πTemp015, ßutcfromtimestamp, nil); πE != nil {
							continue
						}
						if πTemp015, πE = πTemp013.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp010, πE = πg.Eq(πF, πTemp012, πTemp015); πE != nil {
							continue
						}
						πTemp009 = πTemp010
						if πTemp014, πE = πg.IsTrue(πF, πTemp009); πE != nil {
							continue
						}
						if !πTemp014 {
							goto Label13
						}
						if πE = πg.CheckLocal(πF, µnextScheduledAlarm, "nextScheduledAlarm"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp012, πE = πg.GetAttr(πF, µalarm, ßcreated, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GT(πF, µnextScheduledAlarm, πTemp012); πE != nil {
							continue
						}
						πTemp009 = πTemp010
					Label13:
						πTemp008 = πTemp009
					Label12:
						if πTemp011, πE = πg.IsTrue(πF, πTemp008); πE != nil {
							continue
						}
						if πTemp011 {
							goto Label14
						}
						goto Label15
						// line 105: if ((nextScheduledAlarm < datetime.datetime.now()) or
						πF.SetLineno(105)
					Label14:
						// line 107: logging.info(termcolor.colored('4 - Triggering the alarm {} at {} for {}.'.format(
						πF.SetLineno(107)
						πTemp001 = πF.MakeArgs(1)
						πTemp002 = πF.MakeArgs(2)
						πTemp006 = πF.MakeArgs(3)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, µalarm, ßname, nil); πE != nil {
							continue
						}
						πTemp006[0] = πTemp008
						if πTemp008, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp009, ßnow, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp006[1] = πTemp009
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, µalarm, ßalarmTime, nil); πE != nil {
							continue
						}
						πTemp006[2] = πTemp008
						if πTemp008, πE = πg.GetAttr(πF, πg.NewStr("4 - Triggering the alarm {} at {} for {}.").ToObject(), ßformat, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						πTemp002[0] = πTemp009
						πTemp002[1] = ßgreen.ToObject()
						if πTemp008, πE = πg.ResolveGlobal(πF, ßtermcolor); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßcolored, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, πTemp002, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp002)
						πTemp001[0] = πTemp008
						if πTemp008, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßinfo, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πTemp009.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 109: self.manageAlarmTargets(alarm)
						πF.SetLineno(109)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						πTemp001[0] = µalarm
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, µself, ßmanageAlarmTargets, nil); πE != nil {
							continue
						}
						if πTemp009, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						goto Label15
					Label15:
						goto Label11
					Label11:
						goto Label8
					Label8:
						goto Label6
					Label6:
						goto Label1
						goto Label2
					Label2:
						// line 110: threading.Timer(30.0, self.checkAlarms).start()
						πF.SetLineno(110)
						πTemp001 = πF.MakeArgs(2)
						πTemp001[0] = πg.NewFloat(30.0).ToObject()
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µself, ßcheckAlarms, nil); πE != nil {
							continue
						}
						πTemp001[1] = πTemp003
						if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßTimer, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßstart, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßcheckAlarms.ToObject(), πTemp006); πE != nil {
					continue
				}
				// line 112: def manageAlarmTargets(self, alarm):
				πF.SetLineno(112)
				πTemp003 = make([]πg.Param, 2)
				πTemp003[0] = πg.Param{Name: "self", Def: nil}
				πTemp003[1] = πg.Param{Name: "alarm", Def: nil}
				πTemp007 = πg.NewFunction(πg.NewCode("manageAlarmTargets", "main.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µalarm *πg.Object = πArgs[1]; _ = µalarm
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 bool
					_ = πTemp004
					var πTemp005 *πg.Object
					_ = πTemp005
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 113: self.powerSocket(True)
						πF.SetLineno(113)
						πTemp001 = πF.MakeArgs(1)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßpowerSocket, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µalarm, ßsource, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.Eq(πF, πTemp005, ßfile.ToObject()); πE != nil {
							continue
						}
						πTemp002 = πTemp003
						if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp004 {
							goto Label1
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µalarm, ßsource, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.Eq(πF, πTemp005, ßstream.ToObject()); πE != nil {
							continue
						}
						πTemp002 = πTemp003
					Label1:
						if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp004 {
							goto Label2
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µalarm, ßsource, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Eq(πF, πTemp003, ßtts.ToObject()); πE != nil {
							continue
						}
						if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp004 {
							goto Label3
						}
						goto Label4
						// line 114: if alarm.source == 'file' or alarm.source == 'stream':
						πF.SetLineno(114)
					Label2:
						// line 115: self.playFile(alarm)
						πF.SetLineno(115)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						πTemp001[0] = µalarm
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßplayFile, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						goto Label4
						// line 116: elif alarm.source == 'tts':
						πF.SetLineno(116)
					Label3:
						// line 117: self.playTTS(alarm)
						πF.SetLineno(117)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						πTemp001[0] = µalarm
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßplayTTS, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						goto Label4
					Label4:
						// line 118: alarm.lastAlarm = datetime.datetime.now()
						πF.SetLineno(118)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßnow, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µalarm, ßlastAlarm, πTemp002); πE != nil {
							continue
						}
						// line 119: alarm.playing = True
						πF.SetLineno(119)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µalarm, ßplaying, πTemp003); πE != nil {
							continue
						}
						// line 120: alarm.save()
						πF.SetLineno(120)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µalarm, ßsave, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßmanageAlarmTargets.ToObject(), πTemp007); πE != nil {
					continue
				}
				// line 122: def stopAlarm(self):
				πF.SetLineno(122)
				πTemp003 = make([]πg.Param, 1)
				πTemp003[0] = πg.Param{Name: "self", Def: nil}
				πTemp008 = πg.NewFunction(πg.NewCode("stopAlarm", "main.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µalarm *πg.Object = πg.UnboundLocal; _ = µalarm
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 *πg.Object
					_ = πTemp004
					var πTemp005 *πg.Object
					_ = πTemp005
					var πTemp006 *πg.Object
					_ = πTemp006
					var πTemp007 *πg.Object
					_ = πTemp007
					var πTemp008 *πg.Object
					_ = πTemp008
					var πTemp009 bool
					_ = πTemp009
					var πTemp010 *πg.Object
					_ = πTemp010
					var πTemp011 *πg.Object
					_ = πTemp011
					var πTemp012 πg.KWArgs
					_ = πTemp012
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 123: for alarm in Alarm.select().where(Alarm.playing == 1):
						πF.SetLineno(123)
						πTemp001 = πF.MakeArgs(1)
						if πTemp003, πE = πg.ResolveGlobal(πF, ßAlarm); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßplaying, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Eq(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πTemp002, πE = πg.ResolveGlobal(πF, ßAlarm); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßselect, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßwhere, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp003, πE = πg.Iter(πF, πTemp002); πE != nil {
							continue
						}
					Label1:
						if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
								continue
							}
							if !isStop {
								continue
							}
							πE = nil
							πF.RestoreExc(nil, nil)
							goto Label2
						}
						µalarm = πTemp004
						// line 124: alarm.playing = False
						πF.SetLineno(124)
						if πTemp005, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp006}, πTemp005); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µalarm, ßplaying, πTemp006); πE != nil {
							continue
						}
						// line 125: alarm.save()
						πF.SetLineno(125)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µalarm, ßsave, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = πg.NewInt(0).ToObject()
						if πTemp006, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp006.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						πTemp001 = πF.MakeArgs(1)
						πTemp001[0] = πg.NewStr(",").ToObject()
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, µalarm, ßrepeatDays, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßstrip, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp008.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßsplit, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp009, πE = πg.Contains(πF, πTemp006, πTemp007); πE != nil {
							continue
						}
						πTemp005 = πg.GetBool(πTemp009).ToObject()
						if πTemp009, πE = πg.IsTrue(πF, πTemp005); πE != nil {
							continue
						}
						if πTemp009 {
							goto Label3
						}
						goto Label4
						// line 126: if str(0) in alarm.repeatDays.strip().split(','):
						πF.SetLineno(126)
					Label3:
						// line 127: alarm.delete_instance()
						πF.SetLineno(127)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µalarm, ßdelete_instance, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						goto Label4
					Label4:
						// line 128: History.create(
						πF.SetLineno(128)
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µalarm, ßalarmTime, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßnow, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, πTemp007, ßtime, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp010, πE = πg.GetAttr(πF, πTemp008, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp010, ßnow, nil); πE != nil {
							continue
						}
						if πTemp010, πE = πTemp008.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp001 = πF.MakeArgs(2)
						if πTemp008, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp011, πE = πg.GetAttr(πF, πTemp008, ßdate, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp011, ßtoday, nil); πE != nil {
							continue
						}
						if πTemp011, πE = πTemp008.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp011
						if πE = πg.CheckLocal(πF, µalarm, "alarm"); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, µalarm, ßalarmTime, nil); πE != nil {
							continue
						}
						πTemp001[1] = πTemp008
						if πTemp008, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp011, πE = πg.GetAttr(πF, πTemp008, ßdatetime, nil); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp011, ßcombine, nil); πE != nil {
							continue
						}
						if πTemp011, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp006, πE = πg.Sub(πF, πTemp010, πTemp011); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßtotal_seconds, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp008.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp012 = πg.KWArgs{
							{"triggerTime", πTemp005},
							{"stopTime", πTemp007},
							{"duration", πTemp006},
						}
						if πTemp005, πE = πg.ResolveGlobal(πF, ßHistory); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßcreate, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp006.Call(πF, nil, πTemp012); πE != nil {
							continue
						}
						goto Label1
						goto Label2
					Label2:
						// line 134: self.powerSocket(False)
						πF.SetLineno(134)
						πTemp001 = πF.MakeArgs(1)
						if πTemp002, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßpowerSocket, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 135: self.VLCPlayer.stop()
						πF.SetLineno(135)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßVLCPlayer, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßstop, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßstopAlarm.ToObject(), πTemp008); πE != nil {
					continue
				}
				// line 137: def initDB(self):
				πF.SetLineno(137)
				πTemp003 = make([]πg.Param, 1)
				πTemp003[0] = πg.Param{Name: "self", Def: nil}
				πTemp009 = πg.NewFunction(πg.NewCode("initDB", "main.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var πTemp001 *πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 []*πg.Object
					_ = πTemp003
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 138: DB.connect()
						πF.SetLineno(138)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßDB); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßconnect, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						// line 139: Alarm.create_table(True)
						πF.SetLineno(139)
						πTemp003 = πF.MakeArgs(1)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						πTemp003[0] = πTemp001
						if πTemp001, πE = πg.ResolveGlobal(πF, ßAlarm); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcreate_table, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp003)
						// line 140: History.create_table(True)
						πF.SetLineno(140)
						πTemp003 = πF.MakeArgs(1)
						if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						πTemp003[0] = πTemp001
						if πTemp001, πE = πg.ResolveGlobal(πF, ßHistory); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcreate_table, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp003)
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßinitDB.ToObject(), πTemp009); πE != nil {
					continue
				}
				// line 142: def powerSocket(self, state):
				πF.SetLineno(142)
				πTemp003 = make([]πg.Param, 2)
				πTemp003[0] = πg.Param{Name: "self", Def: nil}
				πTemp003[1] = πg.Param{Name: "state", Def: nil}
				πTemp010 = πg.NewFunction(πg.NewCode("powerSocket", "main.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µstate *πg.Object = πArgs[1]; _ = µstate
					var µsocketState *πg.Object = πg.UnboundLocal; _ = µsocketState
					var πTemp001 *πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 bool
					_ = πTemp003
					var πTemp004 []*πg.Object
					_ = πTemp004
					var πTemp005 *πg.Object
					_ = πTemp005
					var πTemp006 *πg.Object
					_ = πTemp006
					var πTemp007 *πg.Object
					_ = πTemp007
					var πTemp008 *πg.Object
					_ = πTemp008
					var πTemp009 *πg.Object
					_ = πTemp009
					var πTemp010 πg.KWArgs
					_ = πTemp010
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 143: socketState = 0
						πF.SetLineno(143)
						µsocketState = πg.NewInt(0).ToObject()
						if πE = πg.CheckLocal(πF, µstate, "state"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Eq(πF, µstate, πTemp002); πE != nil {
							continue
						}
						if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
							continue
						}
						if πTemp003 {
							goto Label1
						}
						goto Label2
						// line 144: if (state == True):
						πF.SetLineno(144)
					Label1:
						// line 145: socketState = 1
						πF.SetLineno(145)
						µsocketState = πg.NewInt(1).ToObject()
						goto Label2
					Label2:
						// line 146: os.system("sudo send433 {code} {id} {state}".format(code=M.Config['SOCKETS']['code'],
						πF.SetLineno(146)
						πTemp004 = πF.MakeArgs(1)
						πTemp001 = ßcode.ToObject()
						πTemp005 = ßSOCKETS.ToObject()
						if πTemp007, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
							continue
						}
						if πTemp008, πE = πg.GetAttr(πF, πTemp007, ßConfig, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp005); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
							continue
						}
						πTemp001 = ßid.ToObject()
						πTemp006 = ßSOCKETS.ToObject()
						if πTemp008, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
							continue
						}
						if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßConfig, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πg.GetItem(πF, πTemp009, πTemp006); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp001); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µsocketState, "socketState"); πE != nil {
							continue
						}
						πTemp010 = πg.KWArgs{
							{"code", πTemp002},
							{"id", πTemp005},
							{"state", µsocketState},
						}
						if πTemp001, πE = πg.GetAttr(πF, πg.NewStr("sudo send433 {code} {id} {state}").ToObject(), ßformat, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp001.Call(πF, nil, πTemp010); πE != nil {
							continue
						}
						πTemp004[0] = πTemp002
						if πTemp001, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsystem, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßpowerSocket.ToObject(), πTemp010); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}).Eval(πF, πF.Globals(), nil, nil)
		if πE != nil {
			continue
		}
		if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			continue
		}
		if πTemp003 == nil {
			πTemp003 = πg.TypeType.ToObject()
		}
		if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("MainMusicAlarm").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßMainMusicAlarm.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 152: class CustomJSONEncoder(flask.json.JSONEncoder):
		πF.SetLineno(152)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp005, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
			continue
		}
		if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßjson, nil); πE != nil {
			continue
		}
		if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßJSONEncoder, nil); πE != nil {
			continue
		}
		πTemp002[0] = πTemp005
		πTemp004 = πg.NewDict()
		if πTemp001, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp001); πE != nil {
			continue
		}
		_, πE = πg.NewCode("CustomJSONEncoder", "main.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp004
			_ = πClass
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 []πg.Param
			_ = πTemp002
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 154: def default(self, obj):
				πF.SetLineno(154)
				πTemp002 = make([]πg.Param, 2)
				πTemp002[0] = πg.Param{Name: "self", Def: nil}
				πTemp002[1] = πg.Param{Name: "obj", Def: nil}
				πTemp001 = πg.NewFunction(πg.NewCode("default", "main.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µobj *πg.Object = πArgs[1]; _ = µobj
					var µmillis *πg.Object = πg.UnboundLocal; _ = µmillis
					var µiterable *πg.Object = πg.UnboundLocal; _ = µiterable
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 bool
					_ = πTemp004
					var πTemp005 *πg.Object
					_ = πTemp005
					var πTemp006 []*πg.Object
					_ = πTemp006
					var πTemp007 *πg.Object
					_ = πTemp007
					var πTemp008 *πg.BaseException
					_ = πTemp008
					var πTemp009 *πg.Traceback
					_ = πTemp009
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 2: goto Label2
						default: panic("unexpected function state")
						}
						// line 155: try:
						πF.SetLineno(155)
						πF.PushCheckpoint(2)
						πTemp001 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
							continue
						}
						πTemp001[0] = µobj
						if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdatetime, nil); πE != nil {
							continue
						}
						πTemp001[1] = πTemp003
						if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πTemp004 {
							goto Label3
						}
						goto Label4
						// line 156: if isinstance(obj, datetime.datetime):
						πF.SetLineno(156)
					Label3:
						if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µobj, ßutcoffset, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.ResolveGlobal(πF, ßNone); πE != nil {
							continue
						}
						πTemp002 = πg.GetBool(πTemp005 != πTemp003).ToObject()
						if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							continue
						}
						if πTemp004 {
							goto Label5
						}
						goto Label6
						// line 157: if obj.utcoffset() is not None:
						πF.SetLineno(157)
					Label5:
						// line 158: obj = obj - obj.utcoffset()
						πF.SetLineno(158)
						if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, µobj, ßutcoffset, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Sub(πF, µobj, πTemp005); πE != nil {
							continue
						}
						µobj = πTemp002
						goto Label6
					Label6:
						// line 159: millis = int(
						πF.SetLineno(159)
						πTemp001 = πF.MakeArgs(1)
						πTemp006 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µobj, ßtimetuple, nil); πE != nil {
							continue
						}
						if πTemp007, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp006[0] = πTemp007
						if πTemp005, πE = πg.ResolveGlobal(πF, ßcalendar); πE != nil {
							continue
						}
						if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßtimegm, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp007.Call(πF, πTemp006, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp006)
						if πTemp003, πE = πg.Mul(πF, πTemp005, πg.NewInt(1000).ToObject()); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
							continue
						}
						if πTemp007, πE = πg.GetAttr(πF, µobj, ßmicrosecond, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πg.Div(πF, πTemp007, πg.NewInt(1000).ToObject()); πE != nil {
							continue
						}
						if πTemp002, πE = πg.Add(πF, πTemp003, πTemp005); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πTemp002, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µmillis = πTemp003
						// line 163: return millis
						πF.SetLineno(163)
						if πE = πg.CheckLocal(πF, µmillis, "millis"); πE != nil {
							continue
						}
						return µmillis, nil
						goto Label4
					Label4:
						πTemp001 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
							continue
						}
						πTemp001[0] = µobj
						if πTemp002, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßtime, nil); πE != nil {
							continue
						}
						πTemp001[1] = πTemp003
						if πTemp002, πE = πg.ResolveGlobal(πF, ßisinstance); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						if πTemp004, πE = πg.IsTrue(πF, πTemp003); πE != nil {
							continue
						}
						if πTemp004 {
							goto Label7
						}
						goto Label8
						// line 164: if isinstance(obj, datetime.time):
						πF.SetLineno(164)
					Label7:
						// line 165: return obj.isoformat()
						πF.SetLineno(165)
						if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µobj, ßisoformat, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						return πTemp003, nil
						goto Label8
					Label8:
						// line 166: iterable = iter(obj)
						πF.SetLineno(166)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
							continue
						}
						πTemp001[0] = µobj
						if πTemp002, πE = πg.ResolveGlobal(πF, ßiter); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µiterable = πTemp003
						πF.PopCheckpoint()
						// line 170: return list(iterable)
						πF.SetLineno(170)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µiterable, "iterable"); πE != nil {
							continue
						}
						πTemp001[0] = µiterable
						if πTemp002, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						return πTemp003, nil
						goto Label1
					Label2:
						πTemp008, πTemp009 = πF.ExcInfo()
						if πTemp002, πE = πg.ResolveGlobal(πF, ßTypeError); πE != nil {
							continue
						}
						if πTemp004, πE = πg.IsInstance(πF, πTemp008.ToObject(), πTemp002); πE != nil {
							continue
						}
						if πTemp004 {
							goto Label9
						}
						πE = πF.Raise(πTemp008.ToObject(), nil, πTemp009.ToObject())
						continue
						// line 167: except TypeError:
						πF.SetLineno(167)
					Label9:
						// line 168: pass
						πF.SetLineno(168)
						πE = nil
						πF.RestoreExc(nil, nil)
						goto Label1
					Label1:
						// line 171: return flask.json.JSONEncoder.default(self, obj)
						πF.SetLineno(171)
						πTemp001 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						πTemp001[0] = µself
						if πE = πg.CheckLocal(πF, µobj, "obj"); πE != nil {
							continue
						}
						πTemp001[1] = µobj
						if πTemp002, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßjson, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, πTemp003, ßJSONEncoder, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßdefault, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						return πTemp002, nil
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßdefault.ToObject(), πTemp001); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}).Eval(πF, πF.Globals(), nil, nil)
		if πE != nil {
			continue
		}
		if πTemp003, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			continue
		}
		if πTemp003 == nil {
			πTemp003 = πg.TypeType.ToObject()
		}
		if πTemp005, πE = πTemp003.Call(πF, []*πg.Object{πg.NewStr("CustomJSONEncoder").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßCustomJSONEncoder.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 173: app = flask.Flask(__name__)
		πF.SetLineno(173)
		πTemp002 = πF.MakeArgs(1)
		if πTemp001, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
			continue
		}
		πTemp002[0] = πTemp001
		if πTemp001, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
			continue
		}
		if πTemp003, πE = πg.GetAttr(πF, πTemp001, ßFlask, nil); πE != nil {
			continue
		}
		if πTemp001, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßapp.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 174: app.json_encoder = CustomJSONEncoder
		πF.SetLineno(174)
		if πTemp001, πE = πg.ResolveGlobal(πF, ßCustomJSONEncoder); πE != nil {
			continue
		}
		if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp001); πE != nil {
			continue
		}
		if πTemp005, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πE = πg.SetAttr(πF, πTemp005, ßjson_encoder, πTemp003); πE != nil {
			continue
		}
		// line 175: M = MainMusicAlarm()
		πF.SetLineno(175)
		if πTemp001, πE = πg.ResolveGlobal(πF, ßMainMusicAlarm); πE != nil {
			continue
		}
		if πTemp003, πE = πTemp001.Call(πF, nil, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßM.ToObject(), πTemp003); πE != nil {
			continue
		}
		// line 179: def stopAlarm():
		πF.SetLineno(179)
		πTemp008 = make([]πg.Param, 0)
		πTemp001 = πg.NewFunction(πg.NewCode("stopAlarm", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 []*πg.Object
			_ = πTemp003
			var πTemp004 *πg.Dict
			_ = πTemp004
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 180: M.stopAlarm()
				πF.SetLineno(180)
				if πTemp001, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstopAlarm, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				// line 181: return flask.jsonify({'success': True})
				πF.SetLineno(181)
				πTemp003 = πF.MakeArgs(1)
				πTemp004 = πg.NewDict()
				if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				if πE = πTemp004.SetItem(πF, ßsuccess.ToObject(), πTemp001); πE != nil {
					continue
				}
				πTemp001 = πTemp004.ToObject()
				πTemp003[0] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßjsonify, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				return πTemp001, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßstopAlarm.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 178: @app.route('/api/v1/stopAlarms')
		πF.SetLineno(178)
		πTemp002 = πF.MakeArgs(1)
		if πTemp003, πE = πg.ResolveGlobal(πF, ßstopAlarm); πE != nil {
			continue
		}
		πTemp002[0] = πTemp003
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/api/v1/stopAlarms").ToObject()
		if πTemp003, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp005, πE = πg.GetAttr(πF, πTemp003, ßroute, nil); πE != nil {
			continue
		}
		if πTemp003, πE = πTemp005.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp005, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßstopAlarm.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 185: def createAlarm_page():
		πF.SetLineno(185)
		πTemp008 = make([]πg.Param, 0)
		πTemp003 = πg.NewFunction(πg.NewCode("createAlarm_page", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πTemp006 *πg.Object
			_ = πTemp006
			var πTemp007 []*πg.Object
			_ = πTemp007
			var πTemp008 *πg.Object
			_ = πTemp008
			var πTemp009 *πg.Object
			_ = πTemp009
			var πTemp010 *πg.Object
			_ = πTemp010
			var πTemp011 πg.KWArgs
			_ = πTemp011
			var πTemp012 *πg.Dict
			_ = πTemp012
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 186: Alarm.create(name=flask.request.json['name'],
				πF.SetLineno(186)
				πTemp001 = ßname.ToObject()
				if πTemp003, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, πTemp004, ßjson, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetItem(πF, πTemp003, πTemp001); πE != nil {
					continue
				}
				πTemp001 = ßsource.ToObject()
				if πTemp004, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßjson, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp001); πE != nil {
					continue
				}
				πTemp001 = ßurl.ToObject()
				if πTemp005, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßjson, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp001); πE != nil {
					continue
				}
				πTemp007 = πF.MakeArgs(2)
				πTemp001 = ßalarmTime.ToObject()
				if πTemp006, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πg.GetAttr(πF, πTemp008, ßjson, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetItem(πF, πTemp006, πTemp001); πE != nil {
					continue
				}
				πTemp007[0] = πTemp005
				πTemp007[1] = πg.NewStr("%H:%M:%S").ToObject()
				if πTemp001, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßdatetime, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßstrptime, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp001.Call(πF, πTemp007, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp007)
				if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßtime, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp001.Call(πF, nil, nil); πE != nil {
					continue
				}
				πTemp001 = ßrepeatDays.ToObject()
				if πTemp008, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp009, πE = πg.GetAttr(πF, πTemp008, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp008, πE = πg.GetAttr(πF, πTemp009, ßjson, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πg.GetItem(πF, πTemp008, πTemp001); πE != nil {
					continue
				}
				πTemp001 = ßactive.ToObject()
				if πTemp009, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp010, πE = πg.GetAttr(πF, πTemp009, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp009, πE = πg.GetAttr(πF, πTemp010, ßjson, nil); πE != nil {
					continue
				}
				if πTemp008, πE = πg.GetItem(πF, πTemp009, πTemp001); πE != nil {
					continue
				}
				πTemp011 = πg.KWArgs{
					{"name", πTemp002},
					{"source", πTemp003},
					{"url", πTemp004},
					{"alarmTime", πTemp005},
					{"repeatDays", πTemp006},
					{"active", πTemp008},
				}
				if πTemp001, πE = πg.ResolveGlobal(πF, ßAlarm); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßcreate, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, πTemp011); πE != nil {
					continue
				}
				// line 193: return flask.jsonify({'success': True})
				πF.SetLineno(193)
				πTemp007 = πF.MakeArgs(1)
				πTemp012 = πg.NewDict()
				if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				if πE = πTemp012.SetItem(πF, ßsuccess.ToObject(), πTemp001); πE != nil {
					continue
				}
				πTemp001 = πTemp012.ToObject()
				πTemp007[0] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßjsonify, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp007)
				return πTemp001, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßcreateAlarm_page.ToObject(), πTemp003); πE != nil {
			continue
		}
		// line 184: @app.route('/api/v1/createAlarm', methods=['POST'])
		πF.SetLineno(184)
		πTemp002 = πF.MakeArgs(1)
		if πTemp005, πE = πg.ResolveGlobal(πF, ßcreateAlarm_page); πE != nil {
			continue
		}
		πTemp002[0] = πTemp005
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/api/v1/createAlarm").ToObject()
		πTemp010 = make([]*πg.Object, 1)
		πTemp010[0] = ßPOST.ToObject()
		πTemp005 = πg.NewList(πTemp010...).ToObject()
		πTemp007 = πg.KWArgs{
			{"methods", πTemp005},
		}
		if πTemp005, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßroute, nil); πE != nil {
			continue
		}
		if πTemp005, πE = πTemp006.Call(πF, πTemp009, πTemp007); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp006, πE = πTemp005.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßcreateAlarm_page.ToObject(), πTemp006); πE != nil {
			continue
		}
		// line 197: def deleteAlarm_page(AlarmId):
		πF.SetLineno(197)
		πTemp008 = make([]πg.Param, 1)
		πTemp008[0] = πg.Param{Name: "AlarmId", Def: nil}
		πTemp005 = πg.NewFunction(πg.NewCode("deleteAlarm_page", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µAlarmId *πg.Object = πArgs[0]; _ = µAlarmId
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 []*πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 198: return flask.jsonify(Alarm.get(Alarm.id == AlarmId).delete_instance())
				πF.SetLineno(198)
				πTemp001 = πF.MakeArgs(1)
				πTemp002 = πF.MakeArgs(1)
				if πTemp004, πE = πg.ResolveGlobal(πF, ßAlarm); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßid, nil); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µAlarmId, "AlarmId"); πE != nil {
					continue
				}
				if πTemp003, πE = πg.Eq(πF, πTemp005, µAlarmId); πE != nil {
					continue
				}
				πTemp002[0] = πTemp003
				if πTemp003, πE = πg.ResolveGlobal(πF, ßAlarm); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßget, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp002)
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdelete_instance, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
					continue
				}
				πTemp001[0] = πTemp003
				if πTemp003, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjsonify, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				return πTemp003, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßdeleteAlarm_page.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 196: @app.route('/api/v1/deleteAlarm/<int:AlarmId>')
		πF.SetLineno(196)
		πTemp002 = πF.MakeArgs(1)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßdeleteAlarm_page); πE != nil {
			continue
		}
		πTemp002[0] = πTemp006
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/api/v1/deleteAlarm/<int:AlarmId>").ToObject()
		if πTemp006, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp011, πE = πg.GetAttr(πF, πTemp006, ßroute, nil); πE != nil {
			continue
		}
		if πTemp006, πE = πTemp011.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp011, πE = πTemp006.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßdeleteAlarm_page.ToObject(), πTemp011); πE != nil {
			continue
		}
		// line 202: def do_testing_page(TestId):
		πF.SetLineno(202)
		πTemp008 = make([]πg.Param, 1)
		πTemp008[0] = πg.Param{Name: "TestId", Def: nil}
		πTemp006 = πg.NewFunction(πg.NewCode("do_testing_page", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µTestId *πg.Object = πArgs[0]; _ = µTestId
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 bool
			_ = πTemp002
			var πTemp003 []*πg.Object
			_ = πTemp003
			var πTemp004 []*πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πTemp006 *πg.Dict
			_ = πTemp006
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				if πE = πg.CheckLocal(πF, µTestId, "TestId"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.Eq(πF, µTestId, πg.NewInt(1).ToObject()); πE != nil {
					continue
				}
				if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
					continue
				}
				if πTemp002 {
					goto Label1
				}
				goto Label2
				// line 203: if TestId == 1:
				πF.SetLineno(203)
			Label1:
				// line 204: M.VLCPlayer.set_media(M.VLCInstance.media_new(
				πF.SetLineno(204)
				πTemp003 = πF.MakeArgs(1)
				πTemp004 = πF.MakeArgs(1)
				πTemp004[0] = πg.NewStr("http://listen.hardbase.fm/tunein-aacplus-pls").ToObject()
				if πTemp001, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßVLCInstance, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßmedia_new, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003[0] = πTemp005
				if πTemp001, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßVLCPlayer, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßset_media, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				// line 206: M.VLCPlayer.play()
				πF.SetLineno(206)
				if πTemp001, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßVLCPlayer, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßplay, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp001.Call(πF, nil, nil); πE != nil {
					continue
				}
				// line 207: M.powerSocket(True)
				πF.SetLineno(207)
				πTemp003 = πF.MakeArgs(1)
				if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				πTemp003[0] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpowerSocket, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				goto Label2
			Label2:
				if πE = πg.CheckLocal(πF, µTestId, "TestId"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.Eq(πF, µTestId, πg.NewInt(2).ToObject()); πE != nil {
					continue
				}
				if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
					continue
				}
				if πTemp002 {
					goto Label3
				}
				goto Label4
				// line 208: if TestId == 2:
				πF.SetLineno(208)
			Label3:
				// line 209: M.powerSocket(False)
				πF.SetLineno(209)
				πTemp003 = πF.MakeArgs(1)
				if πTemp001, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
					continue
				}
				πTemp003[0] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßpowerSocket, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				// line 210: M.VLCPlayer.stop()
				πF.SetLineno(210)
				if πTemp001, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßVLCPlayer, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, πTemp005, ßstop, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp001.Call(πF, nil, nil); πE != nil {
					continue
				}
				goto Label4
			Label4:
				// line 211: return flask.jsonify({'success': True})
				πF.SetLineno(211)
				πTemp003 = πF.MakeArgs(1)
				πTemp006 = πg.NewDict()
				if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				if πE = πTemp006.SetItem(πF, ßsuccess.ToObject(), πTemp001); πE != nil {
					continue
				}
				πTemp001 = πTemp006.ToObject()
				πTemp003[0] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp001, ßjsonify, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp005.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				return πTemp001, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßdo_testing_page.ToObject(), πTemp006); πE != nil {
			continue
		}
		// line 201: @app.route('/api/v1/test/<int:TestId>')
		πF.SetLineno(201)
		πTemp002 = πF.MakeArgs(1)
		if πTemp011, πE = πg.ResolveGlobal(πF, ßdo_testing_page); πE != nil {
			continue
		}
		πTemp002[0] = πTemp011
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/api/v1/test/<int:TestId>").ToObject()
		if πTemp011, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp012, πE = πg.GetAttr(πF, πTemp011, ßroute, nil); πE != nil {
			continue
		}
		if πTemp011, πE = πTemp012.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp012, πE = πTemp011.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßdo_testing_page.ToObject(), πTemp012); πE != nil {
			continue
		}
		// line 215: def get_volume_page(Amount):
		πF.SetLineno(215)
		πTemp008 = make([]πg.Param, 1)
		πTemp008[0] = πg.Param{Name: "Amount", Def: nil}
		πTemp011 = πg.NewFunction(πg.NewCode("get_volume_page", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µAmount *πg.Object = πArgs[0]; _ = µAmount
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 *πg.Dict
			_ = πTemp002
			var πTemp003 []*πg.Object
			_ = πTemp003
			var πTemp004 []*πg.Object
			_ = πTemp004
			var πTemp005 []*πg.Object
			_ = πTemp005
			var πTemp006 *πg.Object
			_ = πTemp006
			var πTemp007 *πg.Object
			_ = πTemp007
			var πTemp008 []*πg.Object
			_ = πTemp008
			var πTemp009 *πg.Object
			_ = πTemp009
			var πTemp010 *πg.Object
			_ = πTemp010
			var πTemp011 πg.KWArgs
			_ = πTemp011
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 216: return flask.jsonify({'data': subprocess.check_output("amixer set Speaker " + str(Amount) + "% | grep % | awk '{print $5}'|sed 's/[^0-9\%]//g' | head -n 1", shell=True).decode("utf-8").rstrip('\n'), 'success': True})
				πF.SetLineno(216)
				πTemp001 = πF.MakeArgs(1)
				πTemp002 = πg.NewDict()
				πTemp003 = πF.MakeArgs(1)
				πTemp003[0] = πg.NewStr("\n").ToObject()
				πTemp004 = πF.MakeArgs(1)
				πTemp004[0] = πg.NewStr("utf-8").ToObject()
				πTemp005 = πF.MakeArgs(1)
				πTemp008 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µAmount, "Amount"); πE != nil {
					continue
				}
				πTemp008[0] = µAmount
				if πTemp009, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
					continue
				}
				if πTemp010, πE = πTemp009.Call(πF, πTemp008, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp008)
				if πTemp007, πE = πg.Add(πF, πg.NewStr("amixer set Speaker ").ToObject(), πTemp010); πE != nil {
					continue
				}
				if πTemp006, πE = πg.Add(πF, πTemp007, πg.NewStr("% | grep % | awk '{print $5}'|sed 's/[^0-9\\%]//g' | head -n 1").ToObject()); πE != nil {
					continue
				}
				πTemp005[0] = πTemp006
				if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				πTemp011 = πg.KWArgs{
					{"shell", πTemp006},
				}
				if πTemp006, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßcheck_output, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp007.Call(πF, πTemp005, πTemp011); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp005)
				if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßdecode, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp007.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßrstrip, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp007.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				if πE = πTemp002.SetItem(πF, ßdata.ToObject(), πTemp006); πE != nil {
					continue
				}
				if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				if πE = πTemp002.SetItem(πF, ßsuccess.ToObject(), πTemp006); πE != nil {
					continue
				}
				πTemp006 = πTemp002.ToObject()
				πTemp001[0] = πTemp006
				if πTemp006, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßjsonify, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp007.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				return πTemp006, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßget_volume_page.ToObject(), πTemp011); πE != nil {
			continue
		}
		// line 214: @app.route('/api/v1/volume/<int:Amount>')
		πF.SetLineno(214)
		πTemp002 = πF.MakeArgs(1)
		if πTemp012, πE = πg.ResolveGlobal(πF, ßget_volume_page); πE != nil {
			continue
		}
		πTemp002[0] = πTemp012
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/api/v1/volume/<int:Amount>").ToObject()
		if πTemp012, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp013, πE = πg.GetAttr(πF, πTemp012, ßroute, nil); πE != nil {
			continue
		}
		if πTemp012, πE = πTemp013.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp013, πE = πTemp012.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßget_volume_page.ToObject(), πTemp013); πE != nil {
			continue
		}
		// line 220: def set_volume_page():
		πF.SetLineno(220)
		πTemp008 = make([]πg.Param, 0)
		πTemp012 = πg.NewFunction(πg.NewCode("set_volume_page", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 *πg.Dict
			_ = πTemp002
			var πTemp003 []*πg.Object
			_ = πTemp003
			var πTemp004 []*πg.Object
			_ = πTemp004
			var πTemp005 []*πg.Object
			_ = πTemp005
			var πTemp006 *πg.Object
			_ = πTemp006
			var πTemp007 πg.KWArgs
			_ = πTemp007
			var πTemp008 *πg.Object
			_ = πTemp008
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 221: return flask.jsonify({'data': subprocess.check_output("amixer get Speaker | grep % | awk '{print $5}'|sed 's/[^0-9\%]//g' | head -n 1", shell=True).decode("utf-8").rstrip('\n'), 'success': True})
				πF.SetLineno(221)
				πTemp001 = πF.MakeArgs(1)
				πTemp002 = πg.NewDict()
				πTemp003 = πF.MakeArgs(1)
				πTemp003[0] = πg.NewStr("\n").ToObject()
				πTemp004 = πF.MakeArgs(1)
				πTemp004[0] = πg.NewStr("utf-8").ToObject()
				πTemp005 = πF.MakeArgs(1)
				πTemp005[0] = πg.NewStr("amixer get Speaker | grep % | awk '{print $5}'|sed 's/[^0-9\\%]//g' | head -n 1").ToObject()
				if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				πTemp007 = πg.KWArgs{
					{"shell", πTemp006},
				}
				if πTemp006, πE = πg.ResolveGlobal(πF, ßsubprocess); πE != nil {
					continue
				}
				if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßcheck_output, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp008.Call(πF, πTemp005, πTemp007); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp005)
				if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßdecode, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp008.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßrstrip, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp008.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				if πE = πTemp002.SetItem(πF, ßdata.ToObject(), πTemp006); πE != nil {
					continue
				}
				if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				if πE = πTemp002.SetItem(πF, ßsuccess.ToObject(), πTemp006); πE != nil {
					continue
				}
				πTemp006 = πTemp002.ToObject()
				πTemp001[0] = πTemp006
				if πTemp006, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp008, πE = πg.GetAttr(πF, πTemp006, ßjsonify, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				return πTemp006, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßset_volume_page.ToObject(), πTemp012); πE != nil {
			continue
		}
		// line 219: @app.route('/api/v1/volume')
		πF.SetLineno(219)
		πTemp002 = πF.MakeArgs(1)
		if πTemp013, πE = πg.ResolveGlobal(πF, ßset_volume_page); πE != nil {
			continue
		}
		πTemp002[0] = πTemp013
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/api/v1/volume").ToObject()
		if πTemp013, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp014, πE = πg.GetAttr(πF, πTemp013, ßroute, nil); πE != nil {
			continue
		}
		if πTemp013, πE = πTemp014.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp014, πE = πTemp013.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßset_volume_page.ToObject(), πTemp014); πE != nil {
			continue
		}
		// line 225: def modifyAlarm_page(AlarmId):
		πF.SetLineno(225)
		πTemp008 = make([]πg.Param, 1)
		πTemp008[0] = πg.Param{Name: "AlarmId", Def: nil}
		πTemp013 = πg.NewFunction(πg.NewCode("modifyAlarm_page", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µAlarmId *πg.Object = πArgs[0]; _ = µAlarmId
			var µentry *πg.Object = πg.UnboundLocal; _ = µentry
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πTemp006 bool
			_ = πTemp006
			var πTemp007 *πg.Object
			_ = πTemp007
			var πTemp008 *πg.Dict
			_ = πTemp008
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 226: entry = Alarm.get(Alarm.id == AlarmId)
				πF.SetLineno(226)
				πTemp001 = πF.MakeArgs(1)
				if πTemp003, πE = πg.ResolveGlobal(πF, ßAlarm); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßid, nil); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µAlarmId, "AlarmId"); πE != nil {
					continue
				}
				if πTemp002, πE = πg.Eq(πF, πTemp004, µAlarmId); πE != nil {
					continue
				}
				πTemp001[0] = πTemp002
				if πTemp002, πE = πg.ResolveGlobal(πF, ßAlarm); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßget, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				µentry = πTemp002
				// line 227: entry.name = flask.request.json[
				πF.SetLineno(227)
				if πTemp004, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßjson, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πg.Contains(πF, πTemp004, ßname.ToObject()); πE != nil {
					continue
				}
				πTemp003 = πg.GetBool(πTemp006).ToObject()
				if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
					continue
				}
				if !πTemp006 {
					goto Label1
				}
				πTemp003 = ßname.ToObject()
				if πTemp005, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßjson, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
					continue
				}
				πTemp002 = πTemp004
				goto Label2
			Label1:
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, µentry, ßname, nil); πE != nil {
					continue
				}
				πTemp002 = πTemp003
			Label2:
				if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πE = πg.SetAttr(πF, µentry, ßname, πTemp003); πE != nil {
					continue
				}
				// line 229: entry.source = flask.request.json[
				πF.SetLineno(229)
				if πTemp004, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßjson, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πg.Contains(πF, πTemp004, ßsource.ToObject()); πE != nil {
					continue
				}
				πTemp003 = πg.GetBool(πTemp006).ToObject()
				if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
					continue
				}
				if !πTemp006 {
					goto Label3
				}
				πTemp003 = ßsource.ToObject()
				if πTemp005, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßjson, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
					continue
				}
				πTemp002 = πTemp004
				goto Label4
			Label3:
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, µentry, ßsource, nil); πE != nil {
					continue
				}
				πTemp002 = πTemp003
			Label4:
				if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πE = πg.SetAttr(πF, µentry, ßsource, πTemp003); πE != nil {
					continue
				}
				// line 231: entry.url = flask.request.json[
				πF.SetLineno(231)
				if πTemp004, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßjson, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πg.Contains(πF, πTemp004, ßurl.ToObject()); πE != nil {
					continue
				}
				πTemp003 = πg.GetBool(πTemp006).ToObject()
				if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
					continue
				}
				if !πTemp006 {
					goto Label5
				}
				πTemp003 = ßurl.ToObject()
				if πTemp005, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßjson, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
					continue
				}
				πTemp002 = πTemp004
				goto Label6
			Label5:
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, µentry, ßurl, nil); πE != nil {
					continue
				}
				πTemp002 = πTemp003
			Label6:
				if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πE = πg.SetAttr(πF, µentry, ßurl, πTemp003); πE != nil {
					continue
				}
				// line 233: entry.alarmTime = flask.request.json[
				πF.SetLineno(233)
				if πTemp004, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßjson, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πg.Contains(πF, πTemp004, ßalarmTime.ToObject()); πE != nil {
					continue
				}
				πTemp003 = πg.GetBool(πTemp006).ToObject()
				if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
					continue
				}
				if !πTemp006 {
					goto Label7
				}
				πTemp003 = ßalarmTime.ToObject()
				if πTemp005, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßjson, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
					continue
				}
				πTemp002 = πTemp004
				goto Label8
			Label7:
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, µentry, ßalarmTime, nil); πE != nil {
					continue
				}
				πTemp002 = πTemp003
			Label8:
				if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πE = πg.SetAttr(πF, µentry, ßalarmTime, πTemp003); πE != nil {
					continue
				}
				// line 235: entry.repeatDays = flask.request.json[
				πF.SetLineno(235)
				if πTemp004, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßjson, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πg.Contains(πF, πTemp004, ßrepeatDays.ToObject()); πE != nil {
					continue
				}
				πTemp003 = πg.GetBool(πTemp006).ToObject()
				if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
					continue
				}
				if !πTemp006 {
					goto Label9
				}
				πTemp003 = ßrepeatDays.ToObject()
				if πTemp005, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßjson, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
					continue
				}
				πTemp002 = πTemp004
				goto Label10
			Label9:
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, µentry, ßrepeatDays, nil); πE != nil {
					continue
				}
				πTemp002 = πTemp003
			Label10:
				if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πE = πg.SetAttr(πF, µentry, ßrepeatDays, πTemp003); πE != nil {
					continue
				}
				// line 237: entry.active = flask.request.json[
				πF.SetLineno(237)
				if πTemp004, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp004, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßjson, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πg.Contains(πF, πTemp004, ßactive.ToObject()); πE != nil {
					continue
				}
				πTemp003 = πg.GetBool(πTemp006).ToObject()
				if πTemp006, πE = πg.IsTrue(πF, πTemp003); πE != nil {
					continue
				}
				if !πTemp006 {
					goto Label11
				}
				πTemp003 = ßactive.ToObject()
				if πTemp005, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp005, ßrequest, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp007, ßjson, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetItem(πF, πTemp005, πTemp003); πE != nil {
					continue
				}
				πTemp002 = πTemp004
				goto Label12
			Label11:
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, µentry, ßactive, nil); πE != nil {
					continue
				}
				πTemp002 = πTemp003
			Label12:
				if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp003}, πTemp002); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πE = πg.SetAttr(πF, µentry, ßactive, πTemp003); πE != nil {
					continue
				}
				// line 239: entry.save()
				πF.SetLineno(239)
				if πE = πg.CheckLocal(πF, µentry, "entry"); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, µentry, ßsave, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				// line 240: return flask.jsonify({'success': True})
				πF.SetLineno(240)
				πTemp001 = πF.MakeArgs(1)
				πTemp008 = πg.NewDict()
				if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				if πE = πTemp008.SetItem(πF, ßsuccess.ToObject(), πTemp002); πE != nil {
					continue
				}
				πTemp002 = πTemp008.ToObject()
				πTemp001[0] = πTemp002
				if πTemp002, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßjsonify, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				return πTemp002, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßmodifyAlarm_page.ToObject(), πTemp013); πE != nil {
			continue
		}
		// line 224: @app.route('/api/v1/modifyAlarm/<int:AlarmId>', methods=['POST'])
		πF.SetLineno(224)
		πTemp002 = πF.MakeArgs(1)
		if πTemp014, πE = πg.ResolveGlobal(πF, ßmodifyAlarm_page); πE != nil {
			continue
		}
		πTemp002[0] = πTemp014
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/api/v1/modifyAlarm/<int:AlarmId>").ToObject()
		πTemp010 = make([]*πg.Object, 1)
		πTemp010[0] = ßPOST.ToObject()
		πTemp014 = πg.NewList(πTemp010...).ToObject()
		πTemp007 = πg.KWArgs{
			{"methods", πTemp014},
		}
		if πTemp014, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp015, πE = πg.GetAttr(πF, πTemp014, ßroute, nil); πE != nil {
			continue
		}
		if πTemp014, πE = πTemp015.Call(πF, πTemp009, πTemp007); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp015, πE = πTemp014.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßmodifyAlarm_page.ToObject(), πTemp015); πE != nil {
			continue
		}
		// line 244: def getAlarms():
		πF.SetLineno(244)
		πTemp008 = make([]πg.Param, 0)
		πTemp014 = πg.NewFunction(πg.NewCode("getAlarms", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 []*πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 245: return flask.jsonify(list(Alarm.select().dicts()))
				πF.SetLineno(245)
				πTemp001 = πF.MakeArgs(1)
				πTemp002 = πF.MakeArgs(1)
				if πTemp003, πE = πg.ResolveGlobal(πF, ßAlarm); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßselect, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdicts, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
					continue
				}
				πTemp002[0] = πTemp003
				if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
					continue
				}
				if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp002)
				πTemp001[0] = πTemp004
				if πTemp003, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjsonify, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				return πTemp003, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßgetAlarms.ToObject(), πTemp014); πE != nil {
			continue
		}
		// line 243: @app.route('/api/v1/getAlarms')
		πF.SetLineno(243)
		πTemp002 = πF.MakeArgs(1)
		if πTemp015, πE = πg.ResolveGlobal(πF, ßgetAlarms); πE != nil {
			continue
		}
		πTemp002[0] = πTemp015
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/api/v1/getAlarms").ToObject()
		if πTemp015, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp016, πE = πg.GetAttr(πF, πTemp015, ßroute, nil); πE != nil {
			continue
		}
		if πTemp015, πE = πTemp016.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp016, πE = πTemp015.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßgetAlarms.ToObject(), πTemp016); πE != nil {
			continue
		}
		// line 249: def getHistory():
		πF.SetLineno(249)
		πTemp008 = make([]πg.Param, 0)
		πTemp015 = πg.NewFunction(πg.NewCode("getHistory", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 []*πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 250: return flask.jsonify(list(History.select().dicts()))
				πF.SetLineno(250)
				πTemp001 = πF.MakeArgs(1)
				πTemp002 = πF.MakeArgs(1)
				if πTemp003, πE = πg.ResolveGlobal(πF, ßHistory); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßselect, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßdicts, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
					continue
				}
				πTemp002[0] = πTemp003
				if πTemp003, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
					continue
				}
				if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp002)
				πTemp001[0] = πTemp004
				if πTemp003, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßjsonify, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				return πTemp003, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßgetHistory.ToObject(), πTemp015); πE != nil {
			continue
		}
		// line 248: @app.route('/api/v1/getHistory')
		πF.SetLineno(248)
		πTemp002 = πF.MakeArgs(1)
		if πTemp016, πE = πg.ResolveGlobal(πF, ßgetHistory); πE != nil {
			continue
		}
		πTemp002[0] = πTemp016
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/api/v1/getHistory").ToObject()
		if πTemp016, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp017, πE = πg.GetAttr(πF, πTemp016, ßroute, nil); πE != nil {
			continue
		}
		if πTemp016, πE = πTemp017.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp017, πE = πTemp016.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßgetHistory.ToObject(), πTemp017); πE != nil {
			continue
		}
		// line 254: def default_serve(path):
		πF.SetLineno(254)
		πTemp008 = make([]πg.Param, 1)
		πTemp008[0] = πg.Param{Name: "path", Def: nil}
		πTemp016 = πg.NewFunction(πg.NewCode("default_serve", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µpath *πg.Object = πArgs[0]; _ = µpath
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 255: return flask.send_from_directory('static', path)
				πF.SetLineno(255)
				πTemp001 = πF.MakeArgs(2)
				πTemp001[0] = ßstatic.ToObject()
				if πE = πg.CheckLocal(πF, µpath, "path"); πE != nil {
					continue
				}
				πTemp001[1] = µpath
				if πTemp002, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsend_from_directory, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				return πTemp002, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßdefault_serve.ToObject(), πTemp016); πE != nil {
			continue
		}
		// line 253: @app.route('/<path:path>')
		πF.SetLineno(253)
		πTemp002 = πF.MakeArgs(1)
		if πTemp017, πE = πg.ResolveGlobal(πF, ßdefault_serve); πE != nil {
			continue
		}
		πTemp002[0] = πTemp017
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/<path:path>").ToObject()
		if πTemp017, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp018, πE = πg.GetAttr(πF, πTemp017, ßroute, nil); πE != nil {
			continue
		}
		if πTemp017, πE = πTemp018.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp018, πE = πTemp017.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßdefault_serve.ToObject(), πTemp018); πE != nil {
			continue
		}
		// line 259: def getFiles_page():
		πF.SetLineno(259)
		πTemp008 = make([]πg.Param, 0)
		πTemp017 = πg.NewFunction(πg.NewCode("getFiles_page", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µfiles *πg.Object = πg.UnboundLocal; _ = µfiles
			var µfilepath *πg.Object = πg.UnboundLocal; _ = µfilepath
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πTemp006 *πg.Object
			_ = πTemp006
			var πTemp007 *πg.Object
			_ = πTemp007
			var πTemp008 πg.KWArgs
			_ = πTemp008
			var πTemp009 *πg.Dict
			_ = πTemp009
			var πTemp010 []*πg.Object
			_ = πTemp010
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 260: files = []
				πF.SetLineno(260)
				πTemp001 = make([]*πg.Object, 0)
				πTemp002 = πg.NewList(πTemp001...).ToObject()
				µfiles = πTemp002
				// line 261: for filepath in glob.iglob(M.Config['ALARM']['musicdirectory'], recursive=True):
				πF.SetLineno(261)
				πTemp001 = πF.MakeArgs(1)
				πTemp002 = ßmusicdirectory.ToObject()
				πTemp004 = ßALARM.ToObject()
				if πTemp006, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßConfig, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetItem(πF, πTemp007, πTemp004); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetItem(πF, πTemp005, πTemp002); πE != nil {
					continue
				}
				πTemp001[0] = πTemp003
				if πTemp002, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				πTemp008 = πg.KWArgs{
					{"recursive", πTemp002},
				}
				if πTemp002, πE = πg.ResolveGlobal(πF, ßglob); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßiglob, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp003.Call(πF, πTemp001, πTemp008); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				if πTemp003, πE = πg.Iter(πF, πTemp002); πE != nil {
					continue
				}
			Label1:
				if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
					isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
					if exc != nil {
						πE = exc
						continue
					}
					if !isStop {
						continue
					}
					πE = nil
					πF.RestoreExc(nil, nil)
					goto Label2
				}
				µfilepath = πTemp004
				// line 262: files.append({
				πF.SetLineno(262)
				πTemp001 = πF.MakeArgs(1)
				πTemp009 = πg.NewDict()
				if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
					continue
				}
				if πE = πTemp009.SetItem(πF, ßpath.ToObject(), µfilepath); πE != nil {
					continue
				}
				πTemp010 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µfilepath, "filepath"); πE != nil {
					continue
				}
				πTemp010[0] = µfilepath
				if πTemp005, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
					continue
				}
				if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßpath, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, πTemp006, ßbasename, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp005.Call(πF, πTemp010, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp010)
				if πE = πTemp009.SetItem(πF, ßfilename.ToObject(), πTemp006); πE != nil {
					continue
				}
				πTemp005 = πTemp009.ToObject()
				πTemp001[0] = πTemp005
				if πE = πg.CheckLocal(πF, µfiles, "files"); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, µfiles, ßappend, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				goto Label1
				goto Label2
			Label2:
				// line 266: return flask.jsonify(files)
				πF.SetLineno(266)
				πTemp001 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µfiles, "files"); πE != nil {
					continue
				}
				πTemp001[0] = µfiles
				if πTemp002, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßjsonify, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				return πTemp002, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßgetFiles_page.ToObject(), πTemp017); πE != nil {
			continue
		}
		// line 258: @app.route('/api/v1/getFiles')
		πF.SetLineno(258)
		πTemp002 = πF.MakeArgs(1)
		if πTemp018, πE = πg.ResolveGlobal(πF, ßgetFiles_page); πE != nil {
			continue
		}
		πTemp002[0] = πTemp018
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/api/v1/getFiles").ToObject()
		if πTemp018, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp019, πE = πg.GetAttr(πF, πTemp018, ßroute, nil); πE != nil {
			continue
		}
		if πTemp018, πE = πTemp019.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp019, πE = πTemp018.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßgetFiles_page.ToObject(), πTemp019); πE != nil {
			continue
		}
		// line 270: def main_page():
		πF.SetLineno(270)
		πTemp008 = make([]πg.Param, 0)
		πTemp018 = πg.NewFunction(πg.NewCode("main_page", "main.py", πTemp008, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 271: return flask.send_from_directory('static', 'index.html')
				πF.SetLineno(271)
				πTemp001 = πF.MakeArgs(2)
				πTemp001[0] = ßstatic.ToObject()
				πTemp001[1] = πg.NewStr("index.html").ToObject()
				if πTemp002, πE = πg.ResolveGlobal(πF, ßflask); πE != nil {
					continue
				}
				if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsend_from_directory, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				return πTemp002, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßmain_page.ToObject(), πTemp018); πE != nil {
			continue
		}
		// line 269: @app.route('/')
		πF.SetLineno(269)
		πTemp002 = πF.MakeArgs(1)
		if πTemp019, πE = πg.ResolveGlobal(πF, ßmain_page); πE != nil {
			continue
		}
		πTemp002[0] = πTemp019
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = πg.NewStr("/").ToObject()
		if πTemp019, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßroute, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		if πTemp020, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πF.Globals().SetItem(πF, ßmain_page.ToObject(), πTemp020); πE != nil {
			continue
		}
		// line 274: class StreamToLogger(object):
		πF.SetLineno(274)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp021, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
			continue
		}
		πTemp002[0] = πTemp021
		πTemp004 = πg.NewDict()
		if πTemp019, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp004.SetItem(πF, ß__module__.ToObject(), πTemp019); πE != nil {
			continue
		}
		_, πE = πg.NewCode("StreamToLogger", "main.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp004
			_ = πClass
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 []πg.Param
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 275: """
				πF.SetLineno(275)
				// line 279: def __init__(self, logger, log_level=logging.DEBUG):
				πF.SetLineno(279)
				πTemp002 = make([]πg.Param, 3)
				πTemp002[0] = πg.Param{Name: "self", Def: nil}
				πTemp002[1] = πg.Param{Name: "logger", Def: nil}
				if πTemp003, πE = πg.ResolveClass(πF, πClass, nil, ßlogging); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßDEBUG, nil); πE != nil {
					continue
				}
				πTemp002[2] = πg.Param{Name: "log_level", Def: πTemp004}
				πTemp001 = πg.NewFunction(πg.NewCode("__init__", "main.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µlogger *πg.Object = πArgs[1]; _ = µlogger
					var µlog_level *πg.Object = πArgs[2]; _ = µlog_level
					var πTemp001 *πg.Object
					_ = πTemp001
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 280: self.logger = logger
						πF.SetLineno(280)
						if πE = πg.CheckLocal(πF, µlogger, "logger"); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlogger); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßlogger, πTemp001); πE != nil {
							continue
						}
						// line 281: self.log_level = log_level
						πF.SetLineno(281)
						if πE = πg.CheckLocal(πF, µlog_level, "log_level"); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µlog_level); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßlog_level, πTemp001); πE != nil {
							continue
						}
						// line 282: self.linebuf = ''
						πF.SetLineno(282)
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, ß.ToObject()); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßlinebuf, πTemp001); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 284: def write(self, buf):
				πF.SetLineno(284)
				πTemp002 = make([]πg.Param, 2)
				πTemp002[0] = πg.Param{Name: "self", Def: nil}
				πTemp002[1] = πg.Param{Name: "buf", Def: nil}
				πTemp003 = πg.NewFunction(πg.NewCode("write", "main.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µbuf *πg.Object = πArgs[1]; _ = µbuf
					var µline *πg.Object = πg.UnboundLocal; _ = µline
					var πTemp001 *πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 []*πg.Object
					_ = πTemp004
					var πTemp005 *πg.Object
					_ = πTemp005
					var πTemp006 *πg.Object
					_ = πTemp006
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 285: for line in buf.rstrip().splitlines():
						πF.SetLineno(285)
						if πE = πg.CheckLocal(πF, µbuf, "buf"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, µbuf, ßrstrip, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßsplitlines, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
							continue
						}
					Label1:
						if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
							isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
							if exc != nil {
								πE = exc
								continue
							}
							if !isStop {
								continue
							}
							πE = nil
							πF.RestoreExc(nil, nil)
							goto Label2
						}
						µline = πTemp003
						// line 286: self.logger.log(self.log_level, line.rstrip())
						πF.SetLineno(286)
						πTemp004 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µself, ßlog_level, nil); πE != nil {
							continue
						}
						πTemp004[0] = πTemp005
						if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µline, ßrstrip, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
							continue
						}
						πTemp004[1] = πTemp006
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp005, πE = πg.GetAttr(πF, µself, ßlogger, nil); πE != nil {
							continue
						}
						if πTemp006, πE = πg.GetAttr(πF, πTemp005, ßlog, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πTemp006.Call(πF, πTemp004, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp004)
						goto Label1
						goto Label2
					Label2:
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßwrite.ToObject(), πTemp003); πE != nil {
					continue
				}
				// line 288: def flush(self):
				πF.SetLineno(288)
				πTemp002 = make([]πg.Param, 1)
				πTemp002[0] = πg.Param{Name: "self", Def: nil}
				πTemp004 = πg.NewFunction(πg.NewCode("flush", "main.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 289: pass
						πF.SetLineno(289)
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßflush.ToObject(), πTemp004); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}).Eval(πF, πF.Globals(), nil, nil)
		if πE != nil {
			continue
		}
		if πTemp020, πE = πTemp004.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			continue
		}
		if πTemp020 == nil {
			πTemp020 = πg.TypeType.ToObject()
		}
		if πTemp021, πE = πTemp020.Call(πF, []*πg.Object{πg.NewStr("StreamToLogger").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp004.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßStreamToLogger.ToObject(), πTemp021); πE != nil {
			continue
		}
		if πTemp020, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
			continue
		}
		if πTemp019, πE = πg.Eq(πF, πTemp020, ß__main__.ToObject()); πE != nil {
			continue
		}
		if πTemp022, πE = πg.IsTrue(πF, πTemp019); πE != nil {
			continue
		}
		if πTemp022 {
			goto Label1
		}
		goto Label2
		// line 291: if (__name__ == "__main__"):
		πF.SetLineno(291)
	Label1:
		// line 292: try:
		πF.SetLineno(292)
		πF.PushCheckpoint(4)
		// line 293: sys.stdout = StreamToLogger(
		πF.SetLineno(293)
		πTemp002 = πF.MakeArgs(2)
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = ßSTDOUT.ToObject()
		if πTemp019, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßgetLogger, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		πTemp002[0] = πTemp019
		if πTemp019, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßINFO, nil); πE != nil {
			continue
		}
		πTemp002[1] = πTemp020
		if πTemp019, πE = πg.ResolveGlobal(πF, ßStreamToLogger); πE != nil {
			continue
		}
		if πTemp020, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πTemp020); πE != nil {
			continue
		}
		if πTemp021, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πE = πg.SetAttr(πF, πTemp021, ßstdout, πTemp019); πE != nil {
			continue
		}
		// line 295: sys.stderr = StreamToLogger(
		πF.SetLineno(295)
		πTemp002 = πF.MakeArgs(2)
		πTemp009 = πF.MakeArgs(1)
		πTemp009[0] = ßSTDERR.ToObject()
		if πTemp019, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßgetLogger, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		πTemp002[0] = πTemp019
		if πTemp019, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßERROR, nil); πE != nil {
			continue
		}
		πTemp002[1] = πTemp020
		if πTemp019, πE = πg.ResolveGlobal(πF, ßStreamToLogger); πE != nil {
			continue
		}
		if πTemp020, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp019}, πTemp020); πE != nil {
			continue
		}
		if πTemp021, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πE = πg.SetAttr(πF, πTemp021, ßstderr, πTemp019); πE != nil {
			continue
		}
		// line 297: sys.setrecursionlimit(100000)
		πF.SetLineno(297)
		πTemp002 = πF.MakeArgs(1)
		πTemp002[0] = πg.NewInt(100000).ToObject()
		if πTemp019, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßsetrecursionlimit, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		// line 298: mainThread = threading.Thread(target=M.checkAlarms)
		πF.SetLineno(298)
		if πTemp019, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßcheckAlarms, nil); πE != nil {
			continue
		}
		πTemp007 = πg.KWArgs{
			{"target", πTemp020},
		}
		if πTemp019, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßThread, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, nil, πTemp007); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßmainThread.ToObject(), πTemp019); πE != nil {
			continue
		}
		// line 299: mainThread.daemon = True
		πF.SetLineno(299)
		if πTemp019, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
			continue
		}
		if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp020}, πTemp019); πE != nil {
			continue
		}
		if πTemp021, πE = πg.ResolveGlobal(πF, ßmainThread); πE != nil {
			continue
		}
		if πE = πg.SetAttr(πF, πTemp021, ßdaemon, πTemp020); πE != nil {
			continue
		}
		// line 300: mainThread.start()
		πF.SetLineno(300)
		if πTemp019, πE = πg.ResolveGlobal(πF, ßmainThread); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßstart, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, nil, nil); πE != nil {
			continue
		}
		// line 301: app.run(host=M.Config['HTTP']['ListenAddr'], port=int(
		πF.SetLineno(301)
		πTemp019 = ßListenAddr.ToObject()
		πTemp021 = ßHTTP.ToObject()
		if πTemp024, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
			continue
		}
		if πTemp025, πE = πg.GetAttr(πF, πTemp024, ßConfig, nil); πE != nil {
			continue
		}
		if πTemp023, πE = πg.GetItem(πF, πTemp025, πTemp021); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetItem(πF, πTemp023, πTemp019); πE != nil {
			continue
		}
		πTemp002 = πF.MakeArgs(1)
		πTemp019 = ßListenPort.ToObject()
		πTemp023 = ßHTTP.ToObject()
		if πTemp025, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
			continue
		}
		if πTemp026, πE = πg.GetAttr(πF, πTemp025, ßConfig, nil); πE != nil {
			continue
		}
		if πTemp024, πE = πg.GetItem(πF, πTemp026, πTemp023); πE != nil {
			continue
		}
		if πTemp021, πE = πg.GetItem(πF, πTemp024, πTemp019); πE != nil {
			continue
		}
		πTemp002[0] = πTemp021
		if πTemp019, πE = πg.ResolveGlobal(πF, ßint); πE != nil {
			continue
		}
		if πTemp021, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		if πTemp019, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
			continue
		}
		if πTemp023, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
			continue
		}
		πTemp007 = πg.KWArgs{
			{"host", πTemp020},
			{"port", πTemp021},
			{"threaded", πTemp019},
			{"debug", πTemp023},
		}
		if πTemp019, πE = πg.ResolveGlobal(πF, ßapp); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßrun, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, nil, πTemp007); πE != nil {
			continue
		}
		πTemp002 = πF.MakeArgs(1)
		πTemp002[0] = πg.NewStr("Flask was interrupted so the programm will terminate!").ToObject()
		if πTemp019, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
			continue
		}
		if πTemp020, πE = πTemp019.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		// line 303: raise Exception(
		πF.SetLineno(303)
		πE = πF.Raise(πTemp020, nil, nil)
		continue
		πF.PopCheckpoint()
		goto Label3
	Label4:
		πTemp027, πTemp028 = πF.ExcInfo()
		if πTemp019, πE = πg.ResolveGlobal(πF, ßException); πE != nil {
			continue
		}
		if πTemp022, πE = πg.IsInstance(πF, πTemp027.ToObject(), πTemp019); πE != nil {
			continue
		}
		if πTemp022 {
			goto Label5
		}
		πE = πF.Raise(πTemp027.ToObject(), nil, πTemp028.ToObject())
		continue
		// line 305: except Exception as e:
		πF.SetLineno(305)
	Label5:
		if πE = πF.Globals().SetItem(πF, ße.ToObject(), πTemp027.ToObject()); πE != nil {
			continue
		}
		// line 306: logging.critical(termcolor.colored('\nThe programm was terminated at {}.\nError: {}. \nStopping services...'.format(
		πF.SetLineno(306)
		πTemp002 = πF.MakeArgs(1)
		πTemp009 = πF.MakeArgs(2)
		πTemp010 = πF.MakeArgs(2)
		if πTemp019, πE = πg.ResolveGlobal(πF, ßdatetime); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßdatetime, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πg.GetAttr(πF, πTemp020, ßnow, nil); πE != nil {
			continue
		}
		if πTemp020, πE = πTemp019.Call(πF, nil, nil); πE != nil {
			continue
		}
		πTemp010[0] = πTemp020
		πTemp029 = πF.MakeArgs(1)
		if πTemp019, πE = πg.ResolveGlobal(πF, ße); πE != nil {
			continue
		}
		πTemp029[0] = πTemp019
		if πTemp019, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
			continue
		}
		if πTemp020, πE = πTemp019.Call(πF, πTemp029, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp029)
		πTemp010[1] = πTemp020
		if πTemp019, πE = πg.GetAttr(πF, πg.NewStr("\nThe programm was terminated at {}.\nError: {}. \nStopping services...").ToObject(), ßformat, nil); πE != nil {
			continue
		}
		if πTemp020, πE = πTemp019.Call(πF, πTemp010, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp010)
		πTemp009[0] = πTemp020
		πTemp009[1] = ßred.ToObject()
		if πTemp019, πE = πg.ResolveGlobal(πF, ßtermcolor); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßcolored, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, πTemp009, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp009)
		πTemp002[0] = πTemp019
		if πTemp019, πE = πg.ResolveGlobal(πF, ßlogging); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßcritical, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		// line 308: M.powerSocket(False)
		πF.SetLineno(308)
		πTemp002 = πF.MakeArgs(1)
		if πTemp019, πE = πg.ResolveGlobal(πF, ßFalse); πE != nil {
			continue
		}
		πTemp002[0] = πTemp019
		if πTemp019, πE = πg.ResolveGlobal(πF, ßM); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßpowerSocket, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		// line 309: sys.exit(os.EX_OK)
		πF.SetLineno(309)
		πTemp002 = πF.MakeArgs(1)
		if πTemp019, πE = πg.ResolveGlobal(πF, ßos); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßEX_OK, nil); πE != nil {
			continue
		}
		πTemp002[0] = πTemp020
		if πTemp019, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
			continue
		}
		if πTemp020, πE = πg.GetAttr(πF, πTemp019, ßexit, nil); πE != nil {
			continue
		}
		if πTemp019, πE = πTemp020.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		πE = nil
		πF.RestoreExc(nil, nil)
		goto Label3
	Label3:
		goto Label2
	Label2:
		return nil, nil
	}
	return nil, πE
}
var Code *πg.Code
func main() {
	Code = πg.NewCode("<module>", "main.py", nil, 0, initModule)
	π_os.Exit(πg.RunMain(Code))
}
