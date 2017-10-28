using System.Threading;

namespace DB_Inserter_Slave
{
    class Program
    {
        static void Main(string[] args)
        {
            Thread t1 = new Thread(new ThreadStart(ThreadInserter));

            t1.Start();

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
