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

}

