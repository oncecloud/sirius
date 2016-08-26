package flow

type CostType int;
type IdType int64;
type CapType int;


type CostModel interface {

  //TaskToUnschedule add edge edge from task to unschedule node
  TaskToUnschedule(taskId IdType) CostType
  //UnscheduleToSin add edge from unschedule to sink node
  UnscheduleToSink(jobId IdType) CostType
  //TaskToResource add edge to resource node
  TaskToResource(taskId IdType, rid IdType) CostType
  //ResourceToResource add edge from resource to resource
  ResourceToResource(src, dst IdType) CostType
  //LeafResourceToSink add edge from resource to sink
  LeafResourceToSink(rid IdType) CostType
  //TaskContinue cost of task continue to run
  TaskContinue(tid IdType) CostType
  //Cost of task preemtion
  TaskPreemtion()

  /*
    Machine and task management
  */
  AddMachine()
  RemoveMachine()
  AddTask(tid IdType)
  RemoveTask()

  DebugInfo() string
}
