 using Microsoft.AspNetCore.Mvc;

 namespace main.Controllers {
     public class ScanCodeController : Controller {

         [Route("scancode/scan")]
         public IActionResult Greet(string username) {
             return Ok("This is the Welcome action method...");
         }
     }
 }