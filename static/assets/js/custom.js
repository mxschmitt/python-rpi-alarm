$(document).ready(function () {
    everyThingIsfineText = "Everything is fine...";
    // Click event handler for stopping the current alarm
    $("#alarm-status > a").click(function (e) {
        if (e.target.text != everyThingIsfineText) {
            $.get("/api/v1/stopAlarms", function (data) {
                $("#alarm-status > a").addClass("ui green label");
                $("#alarm-status > a").text("Success! All alarms are stopped...");
            });
        }

    });
    // Handler for getting the current alarm status (of all)
    $.get("/api/v1/getAlarms", function (data) {
        $("#alarm-status > a").text(everyThingIsfineText);
        $.each(data, function () {
            if (this.playing) {
                $("#alarm-status > a").text("Switch off alarms...");
                $("#alarm-status > a").addClass("ui red label");
                return true
            }
        });
    });
    $("#main-menu > a").each(function () {
        if (this.text == currentPage) {
            $(this).addClass("active");
        }
    });
    $("#addAlarmBTN").click(function () {
        $("#addAlarmModal").modal('show');
        $('#createAlarmForm').form({
            fields: {
                color: {
                    identifier: 'alarm-time',
                    rules: [{
                        type: 'regExp',
                        value: /\d{2}:\d{2}:\d{2}/i,
                    }]
                }
            }
        });

    });
});

function editAlarm(alarmId) {
    console.log(alarmId);
}
