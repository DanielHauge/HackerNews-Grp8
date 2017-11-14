import { Thread } from './thread.model';
import { ThreadDisplay } from './thread-display.model';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { RollbarService } from 'angular-rollbar';

@Injectable()
export class ThreadService {

    constructor(private http: Http, 
        private rollbar: RollbarService) {

    }

    addComment(threadId: number, comment: {  username:string; comment: string; password:string;}) {        
        let data = {"post_title": "", "post_text": comment.comment, "hanesst_id": -1, "post_type": "comment", "post_parent": threadId, "username": comment.username, "pwd_hash": comment.password, "post_url": ""}
        return this.http.post(`http://165.227.151.217:9191/post`, data)
            .toPromise();
    }
    getComments(threadId: number) {
        return this.http.get(`http://165.227.151.217:9191/comments/${threadId}` )
                .toPromise()
                .then(response => response.json().comments as any[]);
    }    
	submitThread(data: {username:string; post_type: string; pwd_hash: string; post_title:string; post_url:string; post_parent:number; hanesst_id:number; post_text:string; }) {
        return this.http.post(`http://165.227.151.217:9191/post`, data)
            .toPromise();
    }

    getThreads(counter:number): Promise<ThreadDisplay[]> {
        return this.http.post('http://165.227.151.217:9191/stories',{"dex": counter,"dex_to": counter+20} )
                .toPromise()
                .then( response => {
                    console.log(response.json());
                    return response.json().stories as ThreadDisplay[];
                }); 
                //.then(response => response.json().data as ThreadDisplay[]);
    }
    upvotePost(data:any): Promise<any> {
        console.log('Upvote post');
        
        return this.http.post('http://165.227.151.217:9191/upvote', data )
                .toPromise()
                .then( response => {
                    console.log(response.json());
                }); 
                //.then(response => response.json().data as ThreadDisplay[]);
    }
      // 3. New method also uses PEOPLE variable
    getThread(id: number) : Promise<Thread> {
    /*    return this.http.get('/app/threads')
        .toPromise()
        .then(response => response.json().data.pop() as Thread);
    }*/
        return this.http.get(`http://165.227.151.217:9191/stories/${id}`)
        .toPromise()
        .then( response => {
            var thread:Thread;
            thread = response.json().thread
            thread.commentamount = response.json().comments.length;
            thread.comments = this.filterComments(response.json().comments);
            console.log(response.json());
            console.log(thread);
            return thread as Thread;
        }); 
       // .then(response => response.json().data as Thread);
    }

    filterComments(comments:any[]){
        var filtered: any[] = [];
        var i = 0;
        for (let comment of comments) {
            comment.vote = false;
            comment.index = i;
            
            filtered.push(comment); // 1, "string", false
            i++;
        }
        return filtered;
    }
}