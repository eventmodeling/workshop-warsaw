using System;
using main.Events;
using Newtonsoft.Json;

namespace main.EventBus {
    internal class EventBus : IEventBus {
        private static string EventFolder = "/events/";

        public EventBus() {

        }

        public void Publish(IEvent @event, string nameOfEvent) {
            var eventJson = JsonConvert.SerializeObject(@event);
            var eventFilePath = GetEventFilePath(nameOfEvent);

            System.IO.File.WriteAllText(eventFilePath, eventJson);
        }

        private string GetEventFilePath(string nameOfEvent) {
            var eventId = GetEventId();
            var nameOfEventWithExtension = $"{nameOfEvent}_{eventId}.json";
            var eventFilePath = System.IO.Path.Combine(EventFolder, nameOfEventWithExtension);
            return eventFilePath;
        }

        private string GetEventId()
        {
           return Guid.NewGuid().ToString();
        }
    }
}