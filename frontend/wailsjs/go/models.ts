export namespace models {
	
	export class Note {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    Title: string;
	    Content: string;
	    ProjectID: number;
	    ObjectiveID: number;
	    GoalID: number;
	    TaskID: number;
	
	    static createFrom(source: any = {}) {
	        return new Note(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.Title = source["Title"];
	        this.Content = source["Content"];
	        this.ProjectID = source["ProjectID"];
	        this.ObjectiveID = source["ObjectiveID"];
	        this.GoalID = source["GoalID"];
	        this.TaskID = source["TaskID"];
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
	export class Goal {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    Name: string;
	    // Go type: time
	    EstimatedEndDate: any;
	    Progress: number;
	    Description: string;
	    ObjectiveID: number;
	    Notes: Note[];
	
	    static createFrom(source: any = {}) {
	        return new Goal(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.Name = source["Name"];
	        this.EstimatedEndDate = this.convertValues(source["EstimatedEndDate"], null);
	        this.Progress = source["Progress"];
	        this.Description = source["Description"];
	        this.ObjectiveID = source["ObjectiveID"];
	        this.Notes = this.convertValues(source["Notes"], Note);
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
	export class GoalUpdateDTO {
	    ID: number;
	    // Go type: time
	    UpdatedAt: any;
	    Name: string;
	    // Go type: time
	    EstimatedEndDate: any;
	    Progress: number;
	    Description: string;
	
	    static createFrom(source: any = {}) {
	        return new GoalUpdateDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.Name = source["Name"];
	        this.EstimatedEndDate = this.convertValues(source["EstimatedEndDate"], null);
	        this.Progress = source["Progress"];
	        this.Description = source["Description"];
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
	
	export class Objective {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    Name: string;
	    // Go type: time
	    EstimatedEndDate: any;
	    Progress: number;
	    Description: string;
	    Overseer: string;
	    ProjectID: number;
	    Goals: Goal[];
	    Notes: Note[];
	
	    static createFrom(source: any = {}) {
	        return new Objective(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.Name = source["Name"];
	        this.EstimatedEndDate = this.convertValues(source["EstimatedEndDate"], null);
	        this.Progress = source["Progress"];
	        this.Description = source["Description"];
	        this.Overseer = source["Overseer"];
	        this.ProjectID = source["ProjectID"];
	        this.Goals = this.convertValues(source["Goals"], Goal);
	        this.Notes = this.convertValues(source["Notes"], Note);
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
	export class ObjectiveUpdateDTO {
	    ID: number;
	    // Go type: time
	    UpdatedAt: any;
	    Name: string;
	    Overseer: string;
	    // Go type: time
	    EstimatedEndDate: any;
	    Progress: number;
	    Description: string;
	
	    static createFrom(source: any = {}) {
	        return new ObjectiveUpdateDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.Name = source["Name"];
	        this.Overseer = source["Overseer"];
	        this.EstimatedEndDate = this.convertValues(source["EstimatedEndDate"], null);
	        this.Progress = source["Progress"];
	        this.Description = source["Description"];
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
	export class Project {
	    ID: number;
	    // Go type: time
	    CreatedAt: any;
	    // Go type: time
	    UpdatedAt: any;
	    // Go type: gorm
	    DeletedAt: any;
	    Name: string;
	    // Go type: time
	    EstimatedEndDate: any;
	    Progress: number;
	    Description: string;
	    Overseer: string;
	    Tags: string;
	    Objectives: Objective[];
	    Notes: Note[];
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.DeletedAt = this.convertValues(source["DeletedAt"], null);
	        this.Name = source["Name"];
	        this.EstimatedEndDate = this.convertValues(source["EstimatedEndDate"], null);
	        this.Progress = source["Progress"];
	        this.Description = source["Description"];
	        this.Overseer = source["Overseer"];
	        this.Tags = source["Tags"];
	        this.Objectives = this.convertValues(source["Objectives"], Objective);
	        this.Notes = this.convertValues(source["Notes"], Note);
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
	export class ProjectUpdateDTO {
	    ID: number;
	    // Go type: time
	    UpdatedAt: any;
	    Name: string;
	    Overseer: string;
	    // Go type: time
	    EstimatedEndDate: any;
	    Progress: number;
	    Description: string;
	    Tags: string;
	
	    static createFrom(source: any = {}) {
	        return new ProjectUpdateDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.UpdatedAt = this.convertValues(source["UpdatedAt"], null);
	        this.Name = source["Name"];
	        this.Overseer = source["Overseer"];
	        this.EstimatedEndDate = this.convertValues(source["EstimatedEndDate"], null);
	        this.Progress = source["Progress"];
	        this.Description = source["Description"];
	        this.Tags = source["Tags"];
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

