var IO = {
    init: function() {
        IO.socket = io.connect("", {transports: ["websocket", "polling"]});
        IO.connectEvents();
    },

    connectEvents: function() {
        IO.socket.on('changed', function(data) {
            IO.deviceChanged(data);
        })
        IO.socket.on('list', function(data) {
            IO.listAllDevices(data);
        });
    },

    deviceChanged: function(data) {
        var device = data || {};
        var list = document.getElementById('list');
        for (var i = 0; i < list.children.length; i++) {
            var item = list.children[i];
            if (item.getAttribute('data-id') == device.ID) {
                var status = IO.convertStatus(device.Status);
                item.children[1].children[0].checked = status[0];
                break;
            }
        }
    },

    listAllDevices: function(data) {
        var devices = data || [];
        var list = document.getElementById('list');
        for (var i = 0; i < devices.length; i++) {
            var device = devices[i];
            if (i >= list.children.length) {
                var li = document.createElement('li');
                li.className = 'item group';
                li.innerHTML = '<p></p><label class="toggle"><input type="checkbox"><div class="track"><div class="handle"></div></div></input></label>';
                list.appendChild(li);
            }
            var item = list.children[i];
            item.children[0].innerHTML = IO.capitalize(device.Name);
            item.setAttribute('data-id', device.ID);
            var status = IO.convertStatus(device.Status);
            item.children[1].children[0].checked = status[0];
            item.children[1].children[0].onclick = IO.sendToggle;
        }
    },

    sendToggle: function() {
        var id = this.parentNode.parentNode.getAttribute('data-id');
        var status = IO.convertStatus(this.checked);
        var data = {
            ID: id,
            Status: status[1]
        };
        IO.socket.emit('toggle', data);
    },

    convertStatus: function(status) {
        if (status === true || (status.toString().toLowerCase().localeCompare('on') === 0)) {
            return [true, 'on'];
        } else {
            return [false, 'off'];
        }
    },

    capitalize: function(string) {
        return string.charAt(0).toUpperCase() + string.slice(1);
    }
};
IO.init();
