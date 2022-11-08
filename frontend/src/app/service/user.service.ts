import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";


import { User } from "./model";
import { environment } from "../../environments/environment";

@Injectable()
export class UserService {
    constructor(private http: HttpClient){}

    private base = environment.ServerUrl;
}