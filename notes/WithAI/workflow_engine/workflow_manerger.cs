using System;
using System.Collections.Generic;
using System.Threading;

public class WorkflowManager
{
    private Queue<ChipWorkflow> workflowQueue = new Queue<ChipWorkflow>();
    private ChipWorkflow currentWorkflow = null;
    private int currentStep = -1;

    public void AddWorkflow(ChipWorkflow workflow)
    {
        workflowQueue.Enqueue(workflow);
        Console.WriteLine("Added workflow for chip {0} to queue.", workflow.ChipID);
    }

    public void ProcessWorkflows()
    {
        while (true)
        {
            if (currentWorkflow == null || currentStep == -1 || currentStep >= currentWorkflow.Workflow.Length)
            {
                // If there is no current workflow or if the current workflow has completed, get the next workflow from the queue
                if (workflowQueue.Count == 0)
                {
                    // If there are no workflows in the queue, wait for some time before checking again
                    Thread.Sleep(1000);
                    continue;
                }
                else
                {
                    currentWorkflow = workflowQueue.Dequeue();
                    currentStep = 0;
                    Console.WriteLine("Starting workflow for chip {0}...", currentWorkflow.ChipID);
                }
            }

            ChipWorkflowStep step = currentWorkflow.Workflow[currentStep];
            Console.WriteLine("Processing step {0} for chip {1} (duration: {2} seconds)...", currentStep + 1, currentWorkflow.ChipID, step.DurationSeconds);
            Thread.Sleep(step.DurationSeconds * 1000);

            currentStep++;
        }
    }
}

public class ChipWorkflow
{
    public int ChipID { get; set; }
    public ChipWorkflowStep[] Workflow { get; set; }
}

public class ChipWorkflowStep
{
    public int DurationSeconds { get; set; }
}

class Program
{
    static void Main(string[] args)
    {
        WorkflowManager workflowManager = new WorkflowManager();

        // Add some sample workflows to the queue
        ChipWorkflow chip1Workflow = new ChipWorkflow { ChipID = 1, Workflow = new ChipWorkflowStep[] { new ChipWorkflowStep { DurationSeconds = 10 }, new ChipWorkflowStep { DurationSeconds = 20 }, new ChipWorkflowStep { DurationSeconds = 10800 }, new ChipWorkflowStep { DurationSeconds = 30 }, new ChipWorkflowStep { DurationSeconds = 10800 }, new ChipWorkflowStep { DurationSeconds = 40 } } };
        ChipWorkflow chip2Workflow = new ChipWorkflow { ChipID = 2, Workflow = new ChipWorkflowStep[] { new ChipWorkflowStep { DurationSeconds = 5 }, new ChipWorkflowStep { DurationSeconds = 10 }, new ChipWorkflowStep { DurationSeconds = 10800 }, new ChipWorkflowStep { DurationSeconds = 15 }, new ChipWorkflowStep { DurationSeconds = 10800 }, new ChipWorkflowStep { DurationSeconds = 20 } } };
        ChipWorkflow chip3Workflow = new ChipWorkflow { ChipID = 3, Workflow = new ChipWorkflowStep[] { new ChipWorkflowStep { DurationSeconds = 2 }, new ChipWorkflowStep { DurationSeconds = 4 }, new ChipWorkflowStep { DurationSeconds = 10800 }, new ChipWorkflowStep { DurationSeconds = 6 }, new ChipWorkflowStep { DurationSeconds = 10800 }, new ChipWorkflowStep { DurationSeconds = 8 } } };
        workflowManager.AddWorkflow(chip1Workflow);
        workflowManager.AddWorkflow(chip2Workflow);
        workflowManager.AddWorkflow(chip3Workflow);

        // Start processing workflows
        workflowManager.ProcessWorkflows();
    }
}
