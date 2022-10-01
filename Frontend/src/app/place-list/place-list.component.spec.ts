import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PlaceListComponent } from './place-list.component';
import * as Rx from 'rxjs';
import { delay } from "rxjs/operators";
import {  async, fakeAsync, tick } from '@angular/core/testing';
import { PlaceFetchService } from '../place-fetch.service';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from "@angular/common/http/testing";
import 'zone.js/dist/zone-testing';


describe('PlaceListComponent', () => {

  let component: PlaceListComponent
  let fixture: ComponentFixture<PlaceListComponent>;
  let plyFetch: any;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PlaceListComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PlaceListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(fixture).toBeTruthy();
  });


//   it('should call getPostDetails and get response as array', fakeAsync(() => {
//     const fixture = TestBed.createComponent(PlaceListComponent);
//     const component = fixture.debugElement.componentInstance;
//     const service = fixture.debugElement.injector.get(PlaceFetchService);
//     let spy_getPosts = spyOn(service,"getPlaces").and.callFake(() => {
//       return Rx.of([{postId : 100}]).pipe(delay(2000));
//     });
//     component.getPlacesList();
//     tick(1000);
//     expect(component.showLoadingIndicator).toEqual(true);
//     tick(1000);
//     expect(component.showLoadingIndicator).toEqual(false);
//     expect(component.places).toEqual([{postId : 100}]);
//   })) 
});
