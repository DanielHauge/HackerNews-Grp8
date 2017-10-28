using System;

namespace DB_Inserter_Slave
{
    public class Comment
    {
        //primary key
        public int ID;
        //forign key to Thread
        public int ThreadID;
        //comments content
        public string Name;
        //forign key to User
        public int UserID;
        //local Karma score
        public int CommentKarma;
        //date of comment creation
        public DateTime Time;
        //school API tracking ID
        public int Han_ID;
        //parent and child relation regarding threads and comments
        public int ParentID;
    }
}
