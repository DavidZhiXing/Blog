using IronPython.Hosting;
using Microsoft.Scripting.Hosting;

// ...

public void ProcessWorkflows()
{
    ScriptEngine engine = Python.CreateEngine();

    while (true)
    {
        if (currentWorkflow == null || currentSubWorkflowIndex == -1 || currentSubWorkflowIndex >= currentWorkflow.SubWorkflows.Length)
        {
            // If there is no current workflow or if the current sub-workflow has completed, get the next workflow from the queue
            if (workflowQueue.Count == 0)
            {
                // If there are no workflows in the queue, wait for some time before checking again
                Thread.Sleep(1000);
                continue;
            }
            else
            {
                currentWorkflow = workflowQueue.Dequeue();
                currentSubWorkflowIndex = 0;
                Console.WriteLine("Starting workflow for chip {0}...", currentWorkflow.ChipID);
            }
        }

        SubWorkflow subWorkflow = currentWorkflow.SubWorkflows[currentSubWorkflowIndex];
        Console.WriteLine("Executing sub-workflow {0} for chip {1}...", currentSubWorkflowIndex + 1, currentWorkflow.ChipID);

        // Execute the Python script for the sub-workflow
        ScriptScope scope = engine.CreateScope();
        ScriptSource source = engine.CreateScriptSourceFromFile(subWorkflow.FilePath);
        source.Execute(scope);

        currentSubWorkflowIndex++;

        if (currentSubWorkflowIndex >= currentWorkflow.SubWorkflows.Length)
        {
            currentWorkflow = null;
            currentSubWorkflowIndex = -1;
            Console.WriteLine("Workflow for chip {0} has completed.", currentWorkflow.ChipID);
        }
        else
        {
            Console.WriteLine("Sub-workflow {0} for chip {1} has completed. Idle time...", currentSubWorkflowIndex, currentWorkflow.ChipID);
            Thread.Sleep(IDLE_TIME_SECONDS * 1000);
        }
    }
}
