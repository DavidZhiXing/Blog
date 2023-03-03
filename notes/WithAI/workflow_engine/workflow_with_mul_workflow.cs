public class WorkflowManager
{
    private Queue<Workflow> workflowQueue = new Queue<Workflow>();
    private Workflow currentWorkflow = null;
    private int currentChipIndex = -1;
    private int currentSubWorkflowIndex = -1;

    private const int IDLE_TIME_SECONDS = 3600;

    public void AddWorkflow(Workflow workflow)
    {
        workflowQueue.Enqueue(workflow);
    }

    public void ProcessWorkflows()
    {
        ScriptEngine engine = Python.CreateEngine();

        while (true)
        {
            if (currentWorkflow == null || currentChipIndex == -1 || currentChipIndex >= currentWorkflow.ChipWorkflows.Count)
            {
                // If there is no current workflow or if the current chip has completed, get the next workflow from the queue
                if (workflowQueue.Count == 0)
                {
                    // If there are no workflows in the queue, wait for some time before checking again
                    Thread.Sleep(1000);
                    continue;
                }
                else
                {
                    currentWorkflow = workflowQueue.Dequeue();
                    currentChipIndex = 0;
                    currentSubWorkflowIndex = -1;
                    Console.WriteLine("Starting workflow {0}...", currentWorkflow.WorkflowID);
                }
            }

            ChipWorkflow chipWorkflow = currentWorkflow.ChipWorkflows[currentChipIndex];

            if (currentSubWorkflowIndex == -1 || currentSubWorkflowIndex >= chipWorkflow.SubWorkflows.Count)
            {
                // If there is no current sub-workflow or if the current sub-workflow has completed, move on to the next sub-workflow
                currentSubWorkflowIndex = 0;
            }

            SubWorkflow subWorkflow = chipWorkflow.SubWorkflows[currentSubWorkflowIndex];
            Console.WriteLine("Executing sub-workflow {0} for chip {1} in workflow {2}...", currentSubWorkflowIndex + 1, chipWorkflow.ChipID, currentWorkflow.WorkflowID);

            // Execute the Python script for the sub-workflow
            ScriptScope scope = engine.CreateScope();
            ScriptSource source = engine.CreateScriptSourceFromFile(subWorkflow.FilePath);
            source.Execute(scope);

            currentSubWorkflowIndex++;

            if (currentSubWorkflowIndex >= chipWorkflow.SubWorkflows.Count)
            {
                currentChipIndex++;
                currentSubWorkflowIndex = -1;
                Console.WriteLine("Chip {0} in workflow {1} has completed.", chipWorkflow.ChipID, currentWorkflow.WorkflowID);
            }
            else
            {
                Console.WriteLine("Sub-workflow {0} for chip {1} in workflow {2} has completed. Idle time...", currentSubWorkflowIndex, chipWorkflow.ChipID, currentWorkflow.WorkflowID);
                Thread.Sleep(IDLE_TIME_SECONDS * 1000);
            }
        }
    }
}


public class Workflow
{
    public string WorkflowID { get; set; }
    public List<ChipWorkflow> ChipWorkflows { get; set; }
}

public class ChipWorkflow
{
    public string ChipID { get; set; }
    public List<SubWorkflow> SubWorkflows { get; set; }
}

public class SubWorkflow
{
    public string FilePath { get; set; }
}
