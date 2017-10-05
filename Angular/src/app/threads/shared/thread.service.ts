import { Thread } from './thread.model';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';

@Injectable()
export class ThreadService {

    constructor(private http: Http) {

    }

    addComment(threadId: number, comment: { name: string; comment: string; }) {
        return this.http.post(`/app/threads/${threadId}/comments`, comment)
            .toPromise();
    }

    getEntries(): Promise<Thread[]> {
        return this.http.get('/app/threads')
                .toPromise()
                .then(response => response.json().data as Thread[]);
    }
      // 3. New method also uses PEOPLE variable
    getThread(id: number) : Promise<Thread> {
    /*    return this.http.get('/app/threads')
        .toPromise()
        .then(response => response.json().data.pop() as Thread);
    }*/
        return this.http.get(`/app/threads/${id}`)
        .toPromise()
        .then(response => response.json().data as Thread);
    }
}