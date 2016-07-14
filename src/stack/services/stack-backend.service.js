(function () {
    'use strict';
    angular.module('app.stack')
        .factory('stackBackend', stackBackend);


    /* @ngInject */
    function stackBackend(gHttp) {
        return {
            createStack: createStack,
            listStacks: listStacks,
            getStack: getStack,
            listStackServices: listStackServices,
            upServiceScale: upServiceScale,
            deleteStack: deleteStack
        };
        
        function createStack(data, form) {
            return gHttp.Resource('stack.stacks').post(data, {form: form});
        }
        
        function listStacks() {
            return gHttp.Resource('stack.stacks').get();
        }
        
        function getStack(stackName) {
            return gHttp.Resource('stack.stack', {stack_name: stackName}).get(); 
        }
        
        function deleteStack(stackName) {
            return gHttp.Resource('stack.stack', {stack_name: stackName}).delete();
        }
        
        function listStackServices(stackName) {
            return gHttp.Resource('stack.services', {stack_name: stackName}).get();
        }
        
        function upServiceScale(stackName, serviceID, scale) {
            return gHttp.Resource('stack.service', {stack_name: stackName, service_id: serviceID}).patch({Scale: scale});
        }
    }
})();