export namespace todo {
	
	export class Todo {
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	    id: string;
	    title: string;
	    completed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Todo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.completed = source["completed"];
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
	export class TodoUpdate {
	    // Go type: time
	    updated_at?: any;
	    id: string;
	    title?: string;
	    completed?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new TodoUpdate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.completed = source["completed"];
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

}

