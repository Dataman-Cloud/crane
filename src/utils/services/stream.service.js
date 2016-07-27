(function () {
 'use strict';
 angular.module('app.utils')
     .factory('stream', stream);

 /* @ngInject */
 function stream(utils) {
     
     var StreamCls = buildStreamCls();
     var token;
     
     return {
         Stream: createStream,
         setToken: setToken
     };
     
     function createStream(urlName, params) {
         return new StreamCls(urlName, params);
     }
     
     function setToken(tokenVal) {
         return token = tokenVal;
     }

     function buildStreamCls() {
         function Stream(urlName, params) {
             this.url = utils.buildFullURL(urlName, params)+"?Authorization="+token;
             this.handlers = {};
             this.errorCallback;
         }
         
         Stream.prototype.addHandler = function (eventType, handler) {
             this.handlers[eventType] = handler;
         }
         
         Stream.prototype.setErrorCallback = function (callback) {
             this.errorCallback = callback;
         }
         
         Stream.prototype.start = function () {
             this.events = new EventSource(this.url);
             angular.forEach(this.handlers, function (handler, eventType) {
                 this.events.addEventListener(eventType, handler);
             }.bind(this));
             
             var self=this;
             this.events.onerror = function (event) {
                 self.stop();
                 console.log('user event stream closed due to error.', event);
                 if (self.errorCallback) {
                     self.errorCallback(event);
                 }
             };
         }
         
         Stream.prototype.stop = function () {
             if (this.events !== undefined) {
                 this.events.close();
                 this.events = undefined;
             }
         }
         
         return Stream;
     }
 }

})();