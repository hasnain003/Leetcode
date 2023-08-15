package main

/*
You have a browser of one tab where you start on the homepage and you can visit another url,
 get back in the history number of steps or move forward in the history number of steps.
 https://leetcode.com/problems/design-browser-history/description/?envType=study-plan-v2&envId=programming-skills
*/
type BrowserHistory struct {
	history []string
	forward []string
}

func BrowserHistoryConstructor(homepage string) BrowserHistory {
	return BrowserHistory{
		history: []string{homepage},
		forward: make([]string, 0),
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.history = append(this.history, url)
	this.forward = this.forward[:0]
}

func (this *BrowserHistory) Back(steps int) string {
	for i := 0; i < steps && len(this.history) > 1; i++ {
		this.forward = append(this.forward, this.history[len(this.history)-1])
		this.history = this.history[:len(this.history)-1]
	}
	return this.history[len(this.history)-1]
}

func (this *BrowserHistory) Forward(steps int) string {
	for i := 0; i < steps && len(this.forward) > 0; i++ {
		this.history = append(this.history, this.forward[len(this.forward)-1])
		this.forward = this.forward[:len(this.forward)-1]
	}
	return this.history[len(this.history)-1]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
