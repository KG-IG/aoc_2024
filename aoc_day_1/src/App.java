import java.io.BufferedReader;
import java.io.FileReader;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class App {
    public static void main(String[] args) throws Exception {
        System.out.println("Hello, World!");
        String inputFilePath = "resources/input.csv";
        
        // init columns
        List<Integer> column1 = new ArrayList<Integer>();
        List<Integer> column2 = new ArrayList<Integer>();

        // read file
        try (BufferedReader br = new BufferedReader(new FileReader(inputFilePath))) {
            String line;
            while ((line = br.readLine()) != null) {
                String[] values = line.split("   ");
                // get first column
                //System.out.println(values[0]);
                column1.add(Integer.parseInt(values[0]));
                // get second column
                column2.add(Integer.parseInt(values[1]));
            }
            
            // sort columns
            Collections.sort(column1);
            Collections.sort(column2);
            
            int result1 = 0;
            int result2 = 0;

            // loop through column
            for(int i = 0; i < column1.size(); i++) {
                result1 = result1 + Math.abs(column1.get(i) - column2.get(i));
                result2 = result2 + Collections.frequency(column2, column1.get(i)) * column1.get(i);
            }
            
            System.out.println("result for first task is: " + result1);
            System.out.println("result for second task is: " + result2);
        }

    }
}
