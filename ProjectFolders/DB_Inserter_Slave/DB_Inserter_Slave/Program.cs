using System.Threading;

namespace DB_Inserter_Slave
{
    class Program
    {
        static void Main(string[] args)
        {
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
