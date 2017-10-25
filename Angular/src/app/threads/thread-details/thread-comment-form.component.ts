import { Component, EventEmitter, Input, Output, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { ThreadService } from '../shared/thread.service';
@Component({
    selector: 'app-thread-comment-form',
    templateUrl: 'thread-comment-form.component.html'
})

export class ThreadCommentFormComponent {
    name: string = "";
    comment: string = "";
    @Input() threadId: number;
    @Output() onCommentAdded = new EventEmitter<{name: string; comment:string;}>();
    @ViewChild('commentForm') commentForm: NgForm;
    
    constructor(private threadService: ThreadService) {

    }
    
    onSubmit(commentForm: NgForm) {
        if (this.commentForm.invalid) return;
        let comment = { name: this.name, comment: this.comment };
        this.threadService.addComment(this.threadId, comment)
            .then(() => {
                this.onCommentAdded.emit(comment);
                    this.commentForm.resetForm();
            });
    }
}