import { Routes } from '@angular/router';
import { GraphComponent } from './graph/graph.component';
import { LibraryComponent } from './library/library.component';
import { LruComponent } from './lru/lru.component';

export const routes: Routes = [
    {path: 'overlap', component: GraphComponent },
    {path:'library', component: LibraryComponent},
    {path:'lru', component: LruComponent},
    {path: '', redirectTo: 'overlap', pathMatch: 'full'},
    {path: '**', redirectTo: 'overlap'}
    
];
