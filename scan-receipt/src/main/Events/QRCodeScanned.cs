namespace main.Events {
    public class QRCodeScanned : IEvent {
        
        public string QRCode { get; }

        private QRCodeScanned(string qRCode) {
            QRCode = qRCode;
        }

        static public QRCodeScanned Create(string qRCode) => new QRCodeScanned(qRCode);

        public string GetName()
        {
            return GetType().BaseType.Name;
        }
    }
}