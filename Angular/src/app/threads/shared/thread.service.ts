import { Thread } from './thread.model';
import { ThreadDisplay } from './thread-display.model';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';

@Injectable()
export class ThreadService {

    constructor(private http: Http) {

    }

    addComment(threadId: number, comment: {  comment: string; }) {
        
        let post = {"post_title": "", "post_text": comment, "hanesst_id": -1, "post_type": "comment", "post_parent": threadId, "username": "onebeerdave", "pwd_hash": "fwozXFe7g0", "post_url": ""}
        return this.http.post(`http://165.227.151.217:9191/${threadId}/post`, comment)
            .toPromise();
    }    
	addThread(threadId: number, comment: { name: string; comment: string; }) {
        return this.http.post(`/app/threads/${threadId}/comments`, comment)
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
      // 3. New method also uses PEOPLE variable
    getThread(id: number) : Promise<Thread> {
    /*    return this.http.get('/app/threads')
        .toPromise()
        .then(response => response.json().data.pop() as Thread);
    }*/
        return this.http.get(`http://165.227.151.217:9191/stories/${id}`)
        .toPromise()
        .then( response => {
            console.log(response.json());
            return response.json().thread as Thread;
        }); 
       // .then(response => response.json().data as Thread);
    }
}