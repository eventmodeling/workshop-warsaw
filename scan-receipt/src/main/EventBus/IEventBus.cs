using main.Events;

namespace main.EventBus
{
    public interface IEventBus
    {
         void Publish(IEvent @event, string nameOfEvent);
    }
}