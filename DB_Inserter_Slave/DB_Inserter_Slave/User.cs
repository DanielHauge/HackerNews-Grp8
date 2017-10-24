namespace DB_Inserter_Slave
{
    public class User
    {
        //primary key
        public int ID;
        //is varchar in database with a max count of 20
        public string Name;
        //total karma count 
        public int KarmaPoints;
        //is varchar in database with a max count of 20
        public string Password;
        //is varchar in database with a max count of 80
        public string Email;
    }
}
