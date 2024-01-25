package rowColumnMajor;

import java.util.Random;

public class IterationOrder {
    private static final int[] SIZE_SERIE = new int[]{
        10   , 20   , 50  ,
        100  , 200  , 500 ,
        1000 , 2000 , 5000,
        10000, 20000
    };

    public static void main(String[] args) {
        for (int size : SIZE_SERIE) {
            int [] arr = generateTestData(size);

            addRowMajor(arr, size);
            addColMajor(arr, size);

            int [][]arr2 = generateTestData2(size, arr);

            addColMajor(arr2, size);
            addColMajor(arr2, size);

            System.out.println();
        }

        System.out.println("------------------------------------------");
        for (int size : SIZE_SERIE) {
            Foo[] arr = generateTestFoos(size);
            addRowMajorFoo(arr, size);

            System.out.println();
        }
    }

    private static Foo[] generateTestFoos(int size) {
        long before= System.nanoTime();
        Foo[] arr = new Foo[size*size];
        Random random = new Random();
        for (int i = 0; i < arr.length; i++) {
            arr[i] = new Foo(random.nextInt(1000), 1.0f);
        }
        System.out.println("generated Foos "+ size + " took: " + micros(before) + " micros ");
        return arr;
    }

    private static int[] generateTestData(int size) {
        int [] arr = new int[size*size];
        Random random = new Random();
        for (int i = 0; i < arr.length; i++) {
            arr[i] = random.nextInt(1000);
        }
        return arr;
    }

    private static int[][] generateTestData2(int size, int[] oldarray) {
        int [][] arr = new int[size][];
        for (int i = 0; i < arr.length; i++) {
            arr[i] = new int[size];
            System.arraycopy(oldarray, i * size, arr[i], 0, arr[i].length);
        }
        return arr;
    }

    private static long addRowMajor(int [] arr, int size) {
        long before= System.nanoTime();
        long sum = 0;
        for (int i = 0; i < size; i++) {
            for (int j = 0; j < size; j++) {
                sum += arr[i*size+j];
            }
        }
        System.out.println("Row major , size " + size +" took " + micros(before) + " micros " + sum);
        return sum;
    }

    private static long addColMajor(int [] arr, int size) {
        long before = System.nanoTime();
        long sum = 0;
        for (int j = 0; j < size; j++) {
            for (int i = 0; i < size; i++) {
                sum += arr[i * size + j];
            }
        }
        System.out.println("Col major , size " + size +" took " + micros(before) + " micros " + sum);
        return sum;
    }

    private static long addRowMajor(int [][] arr, int size) {
        long before= System.nanoTime();
        long sum = 0;
        for (int i = 0; i < size; i++) {
            for (int j = 0; j < size; j++) {
                sum += arr[i][j];
            }
        }
        System.out.println(" Row major Nested , size " + size +" took " + micros(before) + " micros " + sum);
        return sum;
    }

    private static long addColMajor(int [][] arr, int size) {
        long before= System.nanoTime();
        long sum = 0;
        for (int j = 0; j < size; j++) {
            for (int i = 0; i < size; i++) {
                sum += arr[i][j];
            }
        }
        System.out.println(" Row major Nested , size " + size +" took " + micros(before) + " micros " + sum);
        return sum;
    }

    private static void addRowMajorFoo(Foo[] arr, int size) {
        long before= System.nanoTime();
        long sum = 0;
        for (int i = 0; i < size; i++) {
            for (int j = 0; j < size; j++) {
                sum += arr[i * size + j].x;
            }
        }
        System.out.println(" Row major Foos , size " + size +" took " + micros(before) + " micros " + sum);
    }

    private static long micros(long before) {
        return (System.nanoTime() - before) / 1000;
    }

    private static class Foo {
        public int x;
        public float f;

        public Foo(int x, float f) {
            this.x = x;
            this.f = f;
        }
    }
}
