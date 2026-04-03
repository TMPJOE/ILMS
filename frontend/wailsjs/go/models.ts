export namespace models {
	
	export class TaskInput {
	    name: string;
	    desc: string;
	
	    static createFrom(source: any = {}) {
	        return new TaskInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.desc = source["desc"];
	    }
	}
	export class TaskOutput {
	    id: number;
	    status: number;
	    name: string;
	    desc: string;
	    date: string;
	
	    static createFrom(source: any = {}) {
	        return new TaskOutput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.status = source["status"];
	        this.name = source["name"];
	        this.desc = source["desc"];
	        this.date = source["date"];
	    }
	}
	export class TaskResponse {
	    Tasks: TaskOutput[];
	    LastId: number;
	    FirstPage: boolean;
	
	    static createFrom(source: any = {}) {
	        return new TaskResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Tasks = this.convertValues(source["Tasks"], TaskOutput);
	        this.LastId = source["LastId"];
	        this.FirstPage = source["FirstPage"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TaskSwapBack {
	    name: string;
	    desc: string;
	
	    static createFrom(source: any = {}) {
	        return new TaskSwapBack(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.desc = source["desc"];
	    }
	}
	export class TaskUpdate {
	    id: number;
	    status: number;
	    name: string;
	    desc: string;
	
	    static createFrom(source: any = {}) {
	        return new TaskUpdate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.status = source["status"];
	        this.name = source["name"];
	        this.desc = source["desc"];
	    }
	}

}

