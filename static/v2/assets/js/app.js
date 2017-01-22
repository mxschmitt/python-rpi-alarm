const TPLNotFound = {
    template: `
<div class="ui grid stackable centered" style="margin-bottom: 20px;">
    <div class="sixteen wide column">
        <h2>Sorry, but the requested URL is not available!</h2>
    </div>
</div`
}
const TPLHome = {
    template: `
<div class="ui grid stackable" style="margin-bottom: 20px;">
    <div class="sixteen wide column">
        <h2>Main page</h2>
    </div>
    <div class="eight wide column">
        <button class="ui button" style="width: 100%" v-on:click="onPlaybackStart">Testing the playback sound system.</button>
    </div>
    <div class="eight wide column">
        <button class="ui button" style="width: 100%" v-on:click="onPlaybackStop">Stop the playback sound system.</button>
    </div>
</div`,
    methods: {
        onPlaybackStart: function() {
            Vue.http.get('/api/v1/test/1').then((response) => {
                console.log(response)
            }, (response) => {
                console.log(response);
            });
        },
        onPlaybackStop: function() {
            Vue.http.get('/api/v1/test/2').then((response) => {
                console.log(response)
            }, (response) => {
                console.log(response);
            });
        }
    }
}
const TPLAlarms = {
    template: `
<div class="ui grid stackable" style="margin-bottom: 20px;">
    <div class="sixteen wide column">
      <div class="ui inverted dimmer" v-bind:class="{ active: loading }">
    <div class="ui text loader">Loading</div>
  </div>
        <h2>Alarms</h2>
        <div class="ui warning message" v-for="error in errors">
            <i class="close icon"></i>
                <div class="header">Validation error!</div>{{ error }}</div>
        <table class="ui compact celled table">
            <thead>
                <tr>
                    <th>Active</th>
                    <th>Name</th>
                    <th>Source</th>
                    <th>URL</th>
                    <th>Wake-up time</th>
                    <th style="min-width: 180px;">Weekdays</th>
                    <th>Edit</th>
                    <th>Delete</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="alarm in alarms">
                    <td class="collapsing">
                        <div class="ui fitted slider checkbox">
                            <input type="checkbox" value="alarm.active">
                            <label></label>
                        </div>
                    </td>
                    <td>{{ alarm.name }}</td>
                    <td>{{ alarm.source }}</td>
                    <td>{{ alarm.url }}</td>
                    <td>{{ alarm.alarmTime }}</td>
                    <td>{{ alarm.repeatDays }}</td>
                    <td class="selectable">
                        <a href="#">Edit</a>
                    </td>
                    <td class="selectable" v-on:click="onDelete(alarm)">
                       <a style="cursor: pointer;">Delete</a>
                    </td>
                </tr>
                <tr>
                    <td class="collapsing">
                        <div class="ui fitted slider checkbox">
                            <input type="checkbox" v-model="alarm.active">
                            <label></label>
                        </div>
                    </td>
                    <td>
                        <div class="ui input">
                            <input type="text" placeholder="Alarm name" v-model="alarm.name">
                        </div>
                    </td>
                    <td>
                        <select class="ui fluid dropdown" v-model="alarm.source">
                            <option value="">Alarm Source</option>
                            <option value="file">File</option>
                            <option value="stream">Radio</option>
                            <option value="tts">Text to speech</option>
                            <option value="spotify" disabled>Spotify</option>
                            <option value="youtube-dl" disabled>Youtube-DL</option>
                        </select>
                    </td>
                    <td>
                        <div class="ui fluid search selection dropdown" v-if="alarm.source === 'file'" id="alarmFileDropdown">
                            <input type="hidden">
                            <i class="dropdown icon"></i>
                            <div class="default text">Select your track</div>
                            <div class="menu">
                                <div class="item" v-for="track in tracks" :data-value="track.path"><i class="music icon"></i>{{ track.filename }}</div>
                            </div>
                        </div>
                        <div v-else>
                            <div class="ui input">
                                <input type="text" placeholder="Alarm URL" v-model="alarmURLs.other">
                            </div>
                        </div>
                    </td>
                    <td>
                        <div class="ui input" style="max-width: 90px">
                            <input type="text" placeholder="Time" pattern="\d{2}:\d{2}:\d{2}" v-model="alarm.alarmTime">
                        </div>
                    </td>
                    <td>
                        <select multiple="" class="ui fluid dropdown" v-model="alarm.repeatDays">
                            <option value="">Repeat days</option>
                            <option value="0">No repating</option>
                            <option value="1">Monday</option>
                            <option value="2">Tuesday</option>
                            <option value="3">Wednesday</option>
                            <option value="4">Thursday</option>
                            <option value="5">Friday</option>
                            <option value="6">Saturday</option>
                            <option value="7">Sunday</option>
                        </select>
                    </td>
                    <td></td>
                    <td></td>
                </tr>
            </tbody>
            <tfoot class="full-width">
                <tr>
                    <th></th>
                    <th colspan="7">
                        <div class="ui right floated small primary labeled icon button" v-on:click="addAlarm"><i class="add to calendar icon"></i> Add alarm </div>
                        <div class="ui right floated small labeled icon button" v-on:click="reloadTable"><i class="refresh icon"></i> Refresh </div>
                        <div class="ui small button">Approve {{ alarmURLs.file }}</div>
                        <div class="ui small disabled button">Approve All </div>
                    </th>
                </tr>
            </tfoot> 
        </table>
    </div>
    </div`,
    data: function() {
        return {
            loading: false,
            errors: [],
            alarmURLs: {
                other: "Your URL"
            },
            alarms: [],
            tracks: [],
            alarm: {
                name: "",
                url: "",
                source: 'file',
                alarmTime: "",
                repeatDays: [],
                active: true
            },
        }
    },
    methods: {
        reloadTable: function() {
            var vm = this;
            Vue.set(this, 'loading', true);
            Vue.http.get('/api/v1/getAlarms').then((response) => {
                Vue.set(vm, 'alarms', response.body);
                Vue.set(vm, 'loading', false);
            }, (response) => {
                Vue.set(vm, 'loading', false);
                console.log(response);
            });
        },
        loadFileSource: function() {
            var vm = this;
            switch (this.alarm.source) {
                case 'file':
                    Vue.set(this, 'loading', true);
                    Vue.http.get('/api/v1/getFiles').then((response) => {
                        Vue.set(vm, 'tracks', response.body);
                        Vue.set(vm, 'loading', false);
                    }, (response) => {
                        Vue.set(vm, 'loading', false);
                        console.log(response);
                    });
                    break;

                default:
                    break;
            }
        },
        addAlarm: function() {
            var self = this,
                result = this.alarm,
                errors = [];
            switch (this.alarm.source) {
                case 'file':
                    result.url = $('#alarmFileDropdown').dropdown('get value');
                    break;
                default:
                    Vue.set(this, 'alarm.url', this.alarmURLs.default);
                    break;
            }
            // Validation checks
            if (result.name.length < 5) {
                errors.push("The alarm name must be at least 5 chars long.")
            }
            if (result.source == "") {
                errors.push("The alarm source can't be empty.")
            }
            if (result.url == "") {
                errors.push("The alarm url can't be empty.")
            }
            if (!/\d{2}:\d{2}:\d{2}/i.test(result.alarmTime)) {
                errors.push("The wake-up time is not valid.")
            }
            Vue.set(this, 'errors', errors);
            if (errors.length == 0) {
                Vue.http.post('/api/v1/createAlarm', JSON.stringify(result)).then((response) => {
                    if (response.body.success) {
                        self.reloadTable();
                        Vue.set(self, 'alarm.name', "behind Jonas is a beer!");
                    }
                }, (response) => {
                    console.log(response);
                });
            }
        },
        onDelete: function(alarm) {
            var self = this;
            Vue.http.get('/api/v1/deleteAlarm/' + alarm.id).then((response) => {
                console.log(response);
                self.reloadTable();
            }, (response) => {
                console.log(response);
                vm.setAlarmStatus(false);
            });
        }
    },
    mounted: function() {
        this.reloadTable();
        this.loadFileSource();
        $('.ui.dropdown').dropdown();
    },
    watch: {
        'alarm.source': function() {
            this.loadFileSource();
        }
    },
    updated: function() {
        $("#alarmFileDropdown").dropdown();
    }
}

const TPLAbout = {
    template: `
<div class="ui grid stackable" style="margin-bottom: 20px;">
    <div class="eight wide column">
        <h2>About</h2>
        <p>Raspberry Pi alarm clock. Controllable over an smart webinterface which is powerd by vue.js. For the styling we are using the semantic ui framework which provides us great components for usage.</p>
        <h3>Changelog:</h3>
        <ul>
            <li>Added updating alarm status menu</li>
            <li>Fixed footer style issue</li>
            <li>Init project</li>
        </ul>
    </div>
</div`
}

const TPLHistory = {
    template: `
<div class="ui grid stackable" style="margin-bottom: 20px;">
    <div class="eight wide column">
        <h2>History</h2>
    </div>
</div`
}

const router = new VueRouter({
    linkActiveClass: 'active',
    routes: [
        { path: '/', component: TPLHome, name: 'home' },
        { path: '/alarms', component: TPLAlarms, name: 'alarms' },
        { path: '/history', component: TPLHistory, name: 'history' },
        { path: '/about', component: TPLAbout, name: 'about' },
        { path: '/404', component: TPLNotFound, name: '404' },
        { path: '*', redirect: '404' }
    ]
});

var vm = new Vue({
    router,
    el: '#app',
    data: {
        alarmIsNotRunning: true
    },
    methods: {
        setAlarmStatus: function(status) {
            Vue.set(this, 'alarmIsNotRunning', !status);
            return status;
        },
        updateAlarmStatus: function() {
            var vm = this;
            Vue.http.get('/api/v1/getAlarms').then((response) => {
                var status = false;
                response.body.forEach(function(alarm) {
                    if (alarm.playing) {
                        status = true;
                        return true;
                    }
                });
                vm.setAlarmStatus(status);
            }, (response) => {
                console.log(response);
                vm.setAlarmStatus(false);
            });
            setTimeout(this.updateAlarmStatus, 2000);
        },
        stopAlarms: function() {
            Vue.http.get('/api/v1/stopAlarms').then((response) => {
                if (response.body.success) {
                    vm.setAlarmStatus(false);
                } else {
                    vm.setAlarmStatus(true);
                }
            }, (response) => {
                console.log(response);
                vm.setAlarmStatus(true);
            });
            console.log("disabled!");
        }
    },
    mounted: function() {
        this.updateAlarmStatus();
    }
});