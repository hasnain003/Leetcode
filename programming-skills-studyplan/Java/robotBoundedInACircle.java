/*
 * On an infinite plane, a robot initially stands at (0, 0) and faces north. Note that:

The north direction is the positive direction of the y-axis.
The south direction is the negative direction of the y-axis.
The east direction is the positive direction of the x-axis.
The west direction is the negative direction of the x-axis.
The robot can receive one of three instructions:

"G": go straight 1 unit.
"L": turn 90 degrees to the left (i.e., anti-clockwise direction).
"R": turn 90 degrees to the right (i.e., clockwise direction).
The robot performs the instructions given in order, and repeats them forever.
Return true if and only if there exists a circle in the plane such that the robot never leaves the circle.
*/

 /*
  * This Problem can be solved by iterating over the instruction 4 times and if the resultant distance in x and y dirsction is (0,0) then true.
  * This problem can also be solve by traversing only once.
  * And if resultant value in x and y direction is 0,0 then true.
  * And if resultant direction after first traversal is not north than also true.
  * Because in every traversal you traverse the same distance.
  * Suppose you traverse $x and $y first time. Then every time you will travel the same distance.
  * If after first traversal you are facing east dirxn then north -> east -> south -> west -> north.
  * north -> west -> south -> east -> north (when after first traversal you are facing west).
  * north -> south -> north -> south -> north (when after first traversal you are facing south).
  * north -> north -> north -> north -> north (when after first traversal you are pointing north only in this case you won't follow a circle.)
  */
class Solution {
    public boolean isRobotBounded(String instructions) {
        int [][]dirxn = {{0,1},{1,0},{0,-1},{-1,0}};
        int curr = 0,x = 0,y = 0;
        for (int i = 0; i < instructions.length();i++) {
            if (instructions.charAt(i) == 'L') {
                curr = (curr + 3) % 4;
               
            } else if (instructions.charAt(i) == 'R') {
                curr = (curr + 1) % 4;
            } else {
                x += dirxn[curr][0];
                y += dirxn[curr][1];
            }
           
        }
        if (x == 0 && y == 0) {
            return true;
        }
        return curr != 0;
    }
}
