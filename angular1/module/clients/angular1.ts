/// <reference path="../typings/angularjs/angular.d.ts" />
module goTransport {
    "use strict";

    export class Angular1 extends Client{

        constructor(public $q : ng.IQService, public $timeout : ng.ITimeoutService) {
            super($q, $timeout);
        }

        public static getInstance($q : ng.IQService, $timeout : ng.ITimeoutService): Client {
            if(!Angular1.instance)
                Angular1.instance = new Angular1($q, $timeout);
            return Angular1.instance;
        }

    }

    //Attach the above to angular
    "use strict";
    angular
        .module("goTransport", ['bd.sockjs']);
    angular
        .module("goTransport")
        .factory("goTransport", ["$q", "$timeout", Angular1.getInstance]);
}