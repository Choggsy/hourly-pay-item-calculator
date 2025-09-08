# IntelliJ Live Templates

---

## 🛠️ How to Add These Templates

1. Open IntelliJ Settings (`Ctrl + Alt + S` or `Cmd + ,`)
2. Navigate to `Editor → Live Templates`
3. Create a new group (e.g. `GoTemplates`)
4. Add each template with its abbreviation, description, and applicable context (`Go → Statement`)
5. Define variables as shown above
6. Save and use in any Go test file

---

## 🧪 Test Templates

| Template Name   | Abbreviation | Purpose                                                                         | Template Text                                                         | Variables                                           |
|-----------------|--------------|---------------------------------------------------------------------------------|-----------------------------------------------------------------------|-----------------------------------------------------|
| **Subtest**     | `trun`       | Creates a Go subtest using `t.Run(...)` with the cursor placed inside the block | `t.Run("$NAME$", func(t *testing.T) {\n    $END$\n})`                 | `$NAME$`: Subtest name<br>`$END$`: Cursor placement |
| **Errorf**      | `te`         | Generates a formatted error message comparing expected vs actual values         | `if result != expected { t.Errorf("$MSG$", expected, result)}`        | `$MSG$`: "Test failed, expected: %v, got: %v"       |
| **Error Check** | `tee`        | Go test error message scenario for validating expected error output             | `err.Error() != "$ERROR_MSG$" { t.Errorf("Expected error: %v", err)}` | `$ERROR_MSG$`: blank                                |                                                             |                                                     |

