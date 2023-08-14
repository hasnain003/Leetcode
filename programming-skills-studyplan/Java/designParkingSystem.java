class ParkingSystem {
    private int []typeCount;
    public ParkingSystem(int big, int medium, int small) {
        this.typeCount = new int[]{big,medium,small};
    }
    
    public boolean addCar(int carType) {
        if (carType > 3 || carType < 1) {
            return false;
        }
        if (this.typeCount[carType - 1] > 0) {
            this.typeCount[carType - 1]--;
            return true;
        }
        return false;
    }
}

/*
Design a parking system for a parking lot. The parking lot has three kinds of parking spaces:
big, medium, and small, with a fixed number of slots for each size.
https://leetcode.com/problems/design-parking-system/description/?envType=study-plan-v2&envId=programming-skills
*/

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * ParkingSystem obj = new ParkingSystem(big, medium, small);
 * boolean param_1 = obj.addCar(carType);
 */