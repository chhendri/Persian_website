#!/bin/bash

# Initialization of variables
filename=$1;
table_name=$2;
datetime=$(date +"%d%m%y_%H%M");

# Make function to repeat string n times
repeat(){
  var="";
  start=1;
  end=$2;
  for ((i=$start; i<=$end; i++)); do echo -n $1; done
}


echo "Start processing of $filename";

# Create an SQL file
filename=${filename::-4};
new_filename=$filename.sql_$datetime
temp_filename=$filename.sql_temp
cp $filename.csv $temp_filename;

# Get the columns 
columns=$(head -n 1 $temp_filename);
columns=$(sed 's/\t/,/g' <<<"$columns");
len_columns="${columns//[^,]}";
len_columns="${#len_columns}";

# Remove the first line of the file
sed -i '1d' $temp_filename;

cat $temp_filename | while read line; do 
  data=$(sed "s/\t/', '/g" <<<"${line}");
  data=$(sed "s/^/'/g" <<< "${data}");
  data=$(sed "s/$/'/g" <<< "${data}");
  len_data="${data//[^,]}";
  len_data="${#len_data}";

  # Check if the number of data matches the number of fields
  if [ "$len_data" -lt "$len_columns" ]; then
    cols_missing=`expr $len_columns - $len_data`;
    to_add=$(repeat ",''" $cols_missing);
    data2="$data$to_add";
  else
    data2=$data;
  fi;

  # Create the SQL command for each line
  sql_command="insert into $table_name ($columns) values ($data2);";
  echo $sql_command >> $new_filename;
done < "$temp_filename";

# Remove the temporary file
rm -f $temp_filename;

