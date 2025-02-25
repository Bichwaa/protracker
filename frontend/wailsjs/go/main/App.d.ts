// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';

export function CreateGoal(arg1:models.Goal):Promise<void>;

export function CreateObjective(arg1:models.Objective):Promise<void>;

export function CreateProject(arg1:models.Project):Promise<void>;

export function CreateTask(arg1:models.Task):Promise<void>;

export function DeleteGoal(arg1:number):Promise<void>;

export function DeleteObjective(arg1:number):Promise<void>;

export function DeleteProject(arg1:number):Promise<void>;

export function DeleteTask(arg1:number):Promise<void>;

export function FetchGoal(arg1:number):Promise<models.Goal>;

export function FetchGoals():Promise<Array<models.Goal>>;

export function FetchObjective(arg1:number):Promise<models.Objective>;

export function FetchObjectives():Promise<Array<models.Objective>>;

export function FetchProject(arg1:number):Promise<models.Project>;

export function FetchProjects():Promise<Array<models.Project>>;

export function FetchTask(arg1:number):Promise<models.Task>;

export function FetchTasks():Promise<Array<models.Task>>;

export function MarkTaskAsCompleted(arg1:number):Promise<void>;

export function UpdateGoal(arg1:models.GoalUpdateDTO):Promise<void>;

export function UpdateObjective(arg1:models.ObjectiveUpdateDTO):Promise<void>;

export function UpdateProject(arg1:models.ProjectUpdateDTO):Promise<void>;

export function UpdateTask(arg1:models.TaskUpdateDTO):Promise<void>;
