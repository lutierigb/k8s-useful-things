// FROM https://github.com/GoogleCloudPlatform/golang-samples/blob/82da276a923cb9f99534d6c9ee4e52f83839c88a/bigtable/quickstart/main.go
package main

// [START bigtable_hw_imports]
import (
        "context"
        "flag"
        "fmt"
        "log"
        "time"

        "cloud.google.com/go/bigtable"
		"google.golang.org/api/option"
)

// [END bigtable_hw_imports]

// User-provided constants.
const (
        tableName        = "Hello-Bigtable"
        columnFamilyName = "cf1"
        columnName       = "greeting"
)

var greetings = []string{"Hello World!", "Hello Cloud Bigtable!", "Hello golang!"}

// sliceContains reports whether the provided string is present in the given slice of strings.
func sliceContains(list []string, target string) bool {
        for _, s := range list {
                if s == target {
                        return true
                }
        }
        return false
}

func main() {
        project := flag.String("project", "", "The Google Cloud Platform project ID. Required.")
        instance := flag.String("instance", "", "The Google Cloud Bigtable instance ID. Required.")
        flag.Parse()

        for _, f := range []string{"project", "instance"} {
                if flag.Lookup(f).Value.String() == "" {
                        log.Fatalf("The %s flag is required.", f)
                }
        }

        //ctx := context.Background()
        ctx, cancel :=  context.WithTimeout(context.Background(), 1 * time.Second)
        defer cancel()
        // Set up admin client, tables, and column families.
        // NewAdminClient uses Application Default Credentials to authenticate.
        // [START bigtable_hw_connect]
        adminClient, err := bigtable.NewAdminClient(ctx, *project, *instance)
        if err != nil {
                log.Fatalf("Could not create admin client: %v", err)
        }
        log.Println("new admin client done")

        // [END bigtable_hw_connect]

        // [START bigtable_hw_create_table]
        tables, err := adminClient.Tables(ctx)
        if err != nil {
                log.Fatalf("Could not fetch table list: %v", err)
        }

        if !sliceContains(tables, tableName) {
                log.Printf("Creating table %s", tableName)
                if err := adminClient.CreateTable(ctx, tableName); err != nil {
                        log.Fatalf("Could not create table %s: %v", tableName, err)
                }
        }

        tblInfo, err := adminClient.TableInfo(ctx, tableName)
        if err != nil {
                log.Fatalf("Could not read info for table %s: %v", tableName, err)
        }

        if !sliceContains(tblInfo.Families, columnFamilyName) {
                if err := adminClient.CreateColumnFamily(ctx, tableName, columnFamilyName); err != nil {
                        log.Fatalf("Could not create column family %s: %v", columnFamilyName, err)
                }
        }
        // [END bigtable_hw_create_table]

	poolSize := 10
        // Set up Bigtable data operations client.
        // NewClient uses Application Default Credentials to authenticate.
        // [START bigtable_hw_connect_data]
        client, err := bigtable.NewClient(ctx, *project, *instance, option.WithGRPCConnectionPool(poolSize))
        if err != nil {
                log.Fatalf("Could not create data operations client: %v", err)
        }
        // [END bigtable_hw_connect_data]

        // [START bigtable_hw_write_rows]
        tbl := client.Open(tableName)
        muts := make([]*bigtable.Mutation, len(greetings))
        rowKeys := make([]string, len(greetings))

        log.Printf("Writing greeting rows to table")
        for i, greeting := range greetings {
                muts[i] = bigtable.NewMutation()
                muts[i].Set(columnFamilyName, columnName, bigtable.Now(), []byte(greeting))

                // Each row has a unique row key.
                //
                // Note: This example uses sequential numeric IDs for simplicity, but
                // this can result in poor performance in a production application.
                // Since rows are stored in sorted order by key, sequential keys can
                // result in poor distribution of operations across nodes.
                //
                // For more information about how to design a Bigtable schema for the
                // best performance, see the documentation:
                //
                //     https://cloud.google.com/bigtable/docs/schema-design
                rowKeys[i] = fmt.Sprintf("%s%d", columnName, i)
        }

        rowErrs, err := tbl.ApplyBulk(ctx, rowKeys, muts)
        if err != nil {
                log.Fatalf("Could not apply bulk row mutation: %v", err)
        }
        if rowErrs != nil {
                for _, rowErr := range rowErrs {
                        log.Printf("Error writing row: %v", rowErr)
                }
                log.Fatalf("Could not write some rows")
        }
        // [END bigtable_hw_write_rows]

        // [START bigtable_hw_get_by_key]
        log.Printf("Getting a single greeting by row key:")
        row, err := tbl.ReadRow(ctx, rowKeys[0], bigtable.RowFilter(bigtable.ColumnFilter(columnName)))
        if err != nil {
                log.Fatalf("Could not read row with key %s: %v", rowKeys[0], err)
        }
        log.Printf("\t%s = %s\n", rowKeys[0], string(row[columnFamilyName][0].Value))
        // [END bigtable_hw_get_by_key]

        // [START bigtable_hw_scan_all]
        log.Printf("Reading all greeting rows:")
        err = tbl.ReadRows(ctx, bigtable.PrefixRange(columnName), func(row bigtable.Row) bool {
                item := row[columnFamilyName][0]
                log.Printf("\t%s = %s\n", item.Row, string(item.Value))
                return true
        }, bigtable.RowFilter(bigtable.ColumnFilter(columnName)))

        if err = client.Close(); err != nil {
                log.Fatalf("Could not close data operations client: %v", err)
        }
        // [END bigtable_hw_scan_all]

        // [START bigtable_hw_delete_table]
        log.Printf("Deleting the table")
        if err = adminClient.DeleteTable(ctx, tableName); err != nil {
                log.Fatalf("Could not delete table %s: %v", tableName, err)
        }

        if err = adminClient.Close(); err != nil {
                log.Fatalf("Could not close admin client: %v", err)
        }
        // [END bigtable_hw_delete_table]
}