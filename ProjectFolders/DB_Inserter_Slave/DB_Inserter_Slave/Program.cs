using System.Threading;

namespace DB_Inserter_Slave
{
    class Program
    {
        public static string dbusername, dbpassword, dbip, rabbituser, rabbitpassword, rabbitip;
        static void Main(string[] args)
        {
            dbusername = args[0];
            dbpassword = args[1];
            dbip = args[2];
            rabbituser = args[3];
            rabbitpassword = args[4];
            rabbitip = args[5];
            Thread t1 = new Thread(new ThreadStart(ThreadInserter));
            Thread t2 = new Thread(new ThreadStart(UserInserter));
            t1.Start();
            t2.Start();

        }
        private static void ThreadInserter()
        {
            RabbitManager rm = new RabbitManager();
            while (true)
            {
                rm.ReceiveMessage("HNPost");
            }
        }
        private static void UserInserter()
        {
            RabbitManager rm = new RabbitManager();
            while (true)
            {
                rm.ReceiveUser("HNUser");
            }
        }
    }
}
