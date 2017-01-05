$(document).ready(function () {
    everyThingIsfineText = "Everything is fine...";
    // Click event handler for stopping the current alarm
    $("#alarm-status > a").click(function (e) {
        if (e.target.text != everyThingIsfineText) {
            $.get("/api/v1/stopAlarms", function (data) {
                $("#alarm-status > a").removeClass("red");
                $("#alarm-status > a").addClass("ui green label");
                $("#alarm-status > a").text("Success! All alarms are stopped...");
            });
        }

    });
    checkIfAlarmIsRunning();
    setInterval(checkIfAlarmIsRunning, 10000);
    $("#main-menu > a").each(function () {
        if (this.text == currentPage) {
            $(this).addClass("active");
        }
    });
    // Handler for opening the createAlarm modal
    $("#addAlarmBTN").click(function () {
        openAlarmModal();
    });
    $("select[name=alarm-source]").change(function (e) {
        reloadFiles(e.currentTarget.value);
    });
    // Handler for submitting the createAlarm Form
    $("#createAlarmFormBTN").click(function (e) {
        // Error Checking
        if ($('#createAlarmForm').submit().hasClass("success")) {
            var source;
            switch ($("select[name=alarm-source]").val()) {
                case "tts":
                    source = $("input[name=alarm-url]").val();
                    break;
                case "file":
                    source = atob($("select[name=alarm-file]").val().toString());
                    break;
                default:
                    break;
            }
            $.ajax({
                url: '/api/v1/createAlarm',
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                data: JSON.stringify({
                    active: $("input[name=alarm-isactive]").is(':checked'),
                    alarmTime: $("input[name=alarm-time]").val(),
                    name: $("input[name=alarm-name]").val(),
                    source: $("select[name=alarm-source]").val(),
                    repeatDays: $("select[name=alarm-repeatdays]").val().join(),
                    url: source
                })
            }).done(function (data) {
                $("#addAlarmModal").modal('hide');
                reloadAlarmTable();
            });
        }
    });
    $("#reloadTableBTN").click(function () {
        reloadAlarmTable();
    })
});

function openAlarmModal() {
    $("#addAlarmModal").modal('show');
    $('#createAlarmForm').form({
        fields: {
            time: {
                identifier: 'alarm-time',
                rules: [{
                    type: 'regExp',
                    value: /\d{2}:\d{2}:\d{2}/i,
                }, {
                    type: 'exactLength[8]',
                    prompt: 'The alarm time must be exactly 8 characters long.'
                }]
            },
            name: {
                identifier: 'alarm-name',
                rules: [{
                    type: 'empty',
                    prompt: 'Please enter a alarm name.'
                }]
            },
            source: {
                identifier: 'alarm-source',
                rules: [{
                    type: 'empty',
                    prompt: 'Please select a source.'
                }]
            }
        }
    });
    reloadFiles($("select[name=alarm-source]").val() ? $("select[name=alarm-source]").val() : "file");
}

function editAlarm(e, alarmId) {
    console.log("Editing " + alarmId);
}

function reloadAlarmTable() {
    $.get('/api/v1/getAlarms', function (data) {
        $("#alarmsTable > tbody").empty();
        $.each(data, function () {
            $("#alarmsTable > tbody:last-child").append('<tr><td>' + this.name + '</td><td>' + this.source + '</td><td>' + this.url + '</td><td>' + this.alarmTime + '</td><td>' + this.repeatDays + '</td><td>' + this.active + '</td><td class="selectable"><a onclick="deleteAlarm(this, ' + this.id + ')">Delete <i class="trash icon"></i></a></td><td class="selectable"><a onclick="editAlarm(this, ' + this.id + ')">Edit <i class="edit icon"></i></a></td></tr>');
        });
    });
}

function deleteAlarm(e, alarmId) {
    $.get('/api/v1/deleteAlarm/' + alarmId, function (data) {
        $(e).closest('tr').remove();
    });
}

function reloadFiles(type) {
    $('select[name=alarm-file]').children('option:not(:first)').remove();
    $("#alarm-file-field").hide();
    $("#alarm-url-field").hide();
    switch (type) {
        case "file":
            $.getJSON('/api/v1/getFiles', function (data) {
                for (var i = 0; i < data.length; i++) {
                    $('select[name=alarm-file]').append($('<option>', {
                        value: btoa(data[i]),
                        text: data[i].replace(/^.*[\\\/]/, '')
                    }));
                }
                $('select[name=alarm-file] option').eq(1).prop('selected', true);
            });
            $("#alarm-file-field").show();
            break;
        case "tts":

            $("#alarm-url-field").show();
            break;
        default:
            break;
    }
}

function checkIfAlarmIsRunning() {
    // Handler for getting the current alarm status (of all)
    $.get("/api/v1/getAlarms", function (data) {
        $("#alarm-status > a").text(everyThingIsfineText);
        $.each(data, function () {
            if (this.playing) {
                $("#alarm-status > a").text("Switch off alarms...");
                $("#alarm-status > a").removeClass("green");
                $("#alarm-status > a").addClass("ui red label");
                return true
            }
        });
    });
}