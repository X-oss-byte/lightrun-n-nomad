import java.util.*;

public class PrimeMainComplex {

    public static class PrimeChecker {
        public static int zero = 0;
        public int num = 0;
        public PrimeChecker self;

        public PrimeChecker(int i) {
            this.num = i;
            this.self = this;
        }

        public boolean isPrime() {
            if (this.num == 2)
                return true;
            if (this.num < 2 || this.num % 2 == 0)
                return false;
            for (int i = 3; i * i <= this.num; i += 2)
                if (this.num % i == 0)
                    return false;
            return true;
        }

        public void doNothing() {
        }
    }

    public static int cnt = 0;

    public static void main(String[] args) throws InterruptedException{
        int non_static_cnt = 0;
        String message = "Total number of primes: ";
        for (int i = 2; i < Math.pow(10, 9); ++i) {
            PrimeChecker pc = new PrimeChecker(i);
            if (pc.isPrime()) {
                non_static_cnt++;
                cnt++;
            }
            pc.doNothing(); // Fake line to insert breakpoints at
        }
        System.out.println(message + cnt);
    }
}
