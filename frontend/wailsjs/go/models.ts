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

