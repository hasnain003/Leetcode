/*
You have a browser of one tab where you start on the homepage and you can visit another url,
 get back in the history number of steps or move forward in the history number of steps.
 https://leetcode.com/problems/design-browser-history/description/?envType=study-plan-v2&envId=programming-skills
*/

class BrowserHistory {
    private Stack<String> history;
    private Stack<String> future;
    public BrowserHistory(String homepage) {
        this.history = new Stack<>();
        this.future = new Stack<>();
        this.history.push(homepage);
    }
    
    public void visit(String url) {
        this.future = new Stack<>();
        this.history.push(url);
    }
    
    public String back(int steps) {
        for (int i = 0; i < steps && this.history.size() > 1; i++) {
            this.future.push(this.history.pop());
        }
        return this.history.peek();
    }
    
    public String forward(int steps) {
        for (int i = 0; i < steps && this.future.size() > 0; i++) {
            this.history.push(this.future.pop());
        }
        return this.history.peek();
    }
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * BrowserHistory obj = new BrowserHistory(homepage);
 * obj.visit(url);
 * String param_2 = obj.back(steps);
 * String param_3 = obj.forward(steps);
 */