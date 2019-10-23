using main.EventBus;
using main.Events;
using Microsoft.AspNetCore.Mvc;

 namespace main.Controllers {
     
     public class ScanCodeController : Controller {
        private IEventBus _eventBus;

        public ScanCodeController(IEventBus eventBus)
         {
             _eventBus = eventBus;
         }
         
         [Route("scancode/scan/{qRCode}")]
         public IActionResult Scan(string qRCode) {
             var qRCodeScanned = QRCodeScanned.Create(qRCode);
             _eventBus.Publish(qRCodeScanned, "QRCodeScanned");
             return Ok(qRCode);
         }
     }
 }