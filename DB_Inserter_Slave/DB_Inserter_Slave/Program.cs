using System.Threading;

namespace DB_Inserter_Slave
{
    class Program
    {
        static void Main(string[] args)
        {
            Thread t1 = new Thread(new ThreadStart(ThreadInserter));
            Thread t2 = new Thread(new ThreadStart(ThreadInserter));
            Thread t3 = new Thread(new ThreadStart(ThreadInserter));
            Thread t4 = new Thread(new ThreadStart(ThreadInserter));
            Thread t5 = new Thread(new ThreadStart(ThreadInserter));
            t1.Start();
            t2.Start();
            t3.Start();
            t4.Start();
            t5.Start();
        }
        private static void ThreadInserter()
        {
            RabbitManager rm = new RabbitManager();
            while (true)
            {
                rm.ReceiveMessage("HNPost");
            }
        }
    }
}
