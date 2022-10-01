import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DataSharingService {

  constructor() { }
  public isUserLoggedIn: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
  public userRole: BehaviorSubject<string> = new BehaviorSubject<string>("");
  public userEmail: BehaviorSubject<string> = new BehaviorSubject<string>("");
  public userName: BehaviorSubject<string> = new BehaviorSubject<string>("");
}
