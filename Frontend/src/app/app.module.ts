import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSliderModule } from '@angular/material/slider';
import { AngularMaterialModule } from  './material.module';
import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';

import { FlexLayoutModule } from '@angular/flex-layout';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { NavBarComponent } from './nav-bar/nav-bar.component'
import { LogInComponent } from './log-in/log-in.component';
import { RegisterComponent } from './register/register.component';
import { HomePageComponent } from './home-page/home-page.component';
import { AboutUsComponent } from './about-us/about-us.component';
import { HttpClientModule } from '@angular/common/http';
import { PlaceFetchService } from './place-fetch.service';
import { PlaceListComponent } from './place-list/place-list.component';
import { GetUserDataService } from './get-user-data.service';
import { ProfileComponent } from './profile/profile.component';
import { ViewProfileComponent } from './view-profile/view-profile.component';
import { EditProfileComponent } from './edit-profile/edit-profile.component';

import { authInterceptorProviders } from './_helpers/auth.interceptor';
import { AddPackageComponent } from './add-package/add-package.component';
import { DataSharingService } from './services/data-sharing.service';
import { ViewPackagesComponent } from './view-packages/view-packages.component';
import { ViewGuideProfileComponent } from './view-guide-profile/view-guide-profile.component';
import { EditGuideProfileComponent } from './edit-guide-profile/edit-guide-profile.component';
import { AddReviewComponent } from './add-review/add-review.component';
import { ListBookedPackagesComponent } from './list-booked-packages/list-booked-packages.component';


@NgModule({
  declarations: [
    AppComponent,
    NavBarComponent,
    LogInComponent,
    RegisterComponent,
    HomePageComponent,
    AboutUsComponent,
    PlaceListComponent,
    ProfileComponent,
    AddPackageComponent,
    ViewProfileComponent,
    EditProfileComponent,
    ViewPackagesComponent,
    ViewGuideProfileComponent,
    EditGuideProfileComponent,
    AddReviewComponent,
    ListBookedPackagesComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatSliderModule,
    AngularMaterialModule,
    FlexLayoutModule,
    FormsModule, 
    ReactiveFormsModule,
    HttpClientModule
  ],
  providers: [PlaceFetchService,authInterceptorProviders,GetUserDataService,DataSharingService
  ],
  bootstrap: [AppComponent],
  schemas: [CUSTOM_ELEMENTS_SCHEMA]
})
export class AppModule { }
