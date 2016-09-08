(function () {
    'use strict';
    angular.module('app.utils')
        .factory('tty', tty);

    /* @ngInject */
    function tty(utils) {
        var ttyCls = createTTYCls();
        var token;
        return {
            TTY: createTTY,
            setToken: setToken
        };
        
        function setToken(tokenVal) {
            token = tokenVal;
        }

        function createTTY(urlName, params) {
            return new ttyCls(urlName, params);
        }
        
        function createTTYCls() {
           
            function TTY(urlName, params) {
                this.url = utils.buildFullURL(urlName, params, true);
                this._openWS();
            }
            
            TTY.prototype._openWS = function () {
                this.ws = new WebSocket(this.url+"?Authorization="+token);
                var self = this;
                var autoReconnect = -1;

                var term;

                var pingTimer;

                this.ws.onopen = function(event) {
                    pingTimer = setInterval(self.sendPing, 30 * 1000);

                    hterm.defaultStorage = new lib.Storage.Local();

                    term = new hterm.Terminal();

                    term.getPrefs().set("send-encoding", "raw");

                    term.onTerminalReady = function() {
                        var io = term.io.push();

                        io.onVTKeystroke = function(str) {
                            self.ws.send("0" + str);
                        };

                        io.sendString = io.onVTKeystroke;

                        io.onTerminalResize = function(columns, rows) {
                            self.ws.send(
                                "2" + JSON.stringify(
                                    {
                                        columns: columns,
                                        rows: rows,
                                    }
                                )
                            )
                        };

                        term.installKeyboard();
                    };

                    term.decorate(document.getElementById("terminal"));
                };

                this.ws.onmessage = function(event) {
                    var data = event.data.slice(1);
                    switch(event.data[0]) {
                    case '0':
                        term.io.writeUTF8(window.atob(data));
                        break;
                    case '1':
                        // pong
                        break;
                    case '2':
                        term.setWindowTitle(data);
                        break;
                    case '3':
                        var preferences = JSON.parse(data);
                        Object.keys(preferences).forEach(function(key) {
                            console.log("Setting " + key + ": " +  preferences[key]);
                            term.getPrefs().set(key, preferences[key]);
                        });
                        break;
                    case '4':
                        autoReconnect = JSON.parse(data);
                        console.log("Enabling reconnect: " + autoReconnect + " seconds")
                        break;
                    }
                };

                this.ws.onclose = function(event) {
                    if (term) {
                        term.uninstallKeyboard();
                        term.io.showOverlay("容器不存在", null);
                    }
                    clearInterval(pingTimer);
                    if (autoReconnect > 0) {
                        setTimeout(openWs, autoReconnect * 1000);
                    }
                };
            };
            
            TTY.prototype.sendPing = function() {
                if (this.ws) {
                    this.ws.send("1");
                }
            };
            
            TTY.prototype.clear = function() {
                if (this.ws) {
                    this.ws.close();
                }
            };
            return TTY;
        }

    }

})();
