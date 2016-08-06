package flow

type CostType int;
type IdType int64;
type CapType int;


type CostModel interface {
  TaskToUnschedule(taskId IdType)
  UnscheduleToSink(jobId IdType)
  TaskToResource(taskId IdType, rid IdType)
  ResourceToResource(src, dst IdType)
  LeafResourceToSink(rid IdType)
  TaskContinue(tid IdType)
  TaskPreemtion()

  AddMachine()
  RemoveMachine()
  AddTask(tid IdType)
  RemoveTask()

}
