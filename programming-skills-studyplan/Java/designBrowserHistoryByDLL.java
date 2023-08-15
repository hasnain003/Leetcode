/*
You have a browser of one tab where you start on the homepage and you can visit another url,
 get back in the history number of steps or move forward in the history number of steps.
 https://leetcode.com/problems/design-browser-history/description/?envType=study-plan-v2&envId=programming-skills
*/

class DLL {
    DLL prev,next;
    String data;
    public DLL(String url) {
        this.data = url;
        this.prev = null;
        this.next = null;
    }
}
class BrowserHistory {
    private DLL head;
    private DLL current;
    public BrowserHistory(String homepage) {
        head = new DLL(homepage);
        current = head;
    }
    
    public void visit(String url) {
        DLL newNode = new DLL(url);
        newNode.prev = current;
        current.next = newNode;
        current = newNode;
    }
    
    public String back(int steps) {
        for (int i = 0; i < steps && current.prev != null;i++) {
            current = current.prev;
        }
        return current.data;
    }
    
    public String forward(int steps) {
        for (int i = 0; i < steps && current.next != null;i++) {
            current = current.next;
        }
        return current.data;
    }
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * BrowserHistory obj = new BrowserHistory(homepage);
 * obj.visit(url);
 * String param_2 = obj.back(steps);
 * String param_3 = obj.forward(steps);
 */