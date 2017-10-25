import { InMemoryDbService } from 'angular-in-memory-web-api';

declare var file: any;

export class InMemoryThreadService implements InMemoryDbService {
    createDb() {
        let threads = [
            {
                id: 1,
 				username: 'John Dohe',
                post_title: 'Colony: A platform for open organizations',
                post_type: 'story',
                post_text: 'text.......',
                post_parent: 'parent',
                comments: [
                    {
                        id: 1,
                        name: 'Jane Smith',
                        comment: 'This is my favorite! I love it!'
                    }
                ]
            },
            {
                id: 2,
				username: 'Anna',
                post_title: 'Hacker News Clone Using GraphQL and React',
                post_type: 'story',
                post_text: 'text....',
                post_parent: 'parent',
                comments: [
                    {
                        id: 2,
                        name: 'Kyle Jones',
                        comment: 'Nice!'
                    },
                    {
                        id: 3,
                        name: 'Alecia Clark',
                        comment: 'All the greens make this amazing.'
                    }
                ]
            },
            {
                id: 3,
                title: '',
				username: 'Niels',
                post_title: 'Whom the Gods Would Destroy, They First Give Real-Time Analytics (2013)',
                post_type: 'story',
                post_text: 'text...',
                post_parent: 'parent',
                comments:[
                    {
                        id: 2,
                        name: 'Kyle Jones',
                        comment: 'Nice!'
                    },
                    {
                        id: 3,
                        name: 'Alecia Clark',
                        comment: 'All the greens make this amazing.'
                    }
                ]
            },
            {
                id: 4,
 				username: 'username',
                post_title: 'Google CEO Appeases Publishers with Subscriptions',
                post_type: 'story',
                post_text: 'text.............',
                post_parent: 'parent',			
                comments: [
                    {
                        id: 4,
                        name: 'Steve Johnson',
                        comment: 'It looks like trouble is on the way.'
                    },
                    {
                        id: 5,
                        name: 'Becky M',
                        comment: 'I imagine this was a shot of a storm that just passed.'
                    }
                ]
            },
            {
                id: 5,
				username: 'username',
                post_title: 'IMF Head Foresees the End of Banking and the Triumph of Cryptocurrency',
                post_type: 'story',
                post_text: 'text',
                post_parent: 'parent',				
                comments: [
                    {
                        id: 6,
                        name: 'Lisa Frank',
                        comment: 'Beautiful!'
                    }
                ]
            }
        ];
        return { threads };
    }
}