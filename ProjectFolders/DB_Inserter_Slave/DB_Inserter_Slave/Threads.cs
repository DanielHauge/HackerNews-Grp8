using System;

namespace DB_Inserter_Slave
{
    public class Threads
    {
        //primary key
        public int ID;
        //is varchar in database with a max count of 5000
        public string Name;
        //forign key directed at User
        public int UserID;
        //date of thread creation
        public DateTime Time;
        //school API tracking ID
        public int Han_ID;
        //is varchar in database with a max count of 200
        public string Post_URL;
    }
}
