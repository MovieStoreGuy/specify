Specify
=======

This library allows for abstracted business for validating conditions are met using the [Specification Pattern]. 

The benefits of using this library are:

- Conditional Logic can be abstracted into unit testable methods.
- Abstractions provide a deeper insight and intent into what is being checked.
- Conditions can be chained together to form higher order rules while remaining testable.
- Increases the reason about the business logic

Since conditions can be appended to and chained together, it means that logic implementation is not repeated
and can be easily refactored in future updates to the project.

## Usage

An example of where _Specify_ can shine is when working with various external packages or systems
that may have the ability to express conditional statements. 
The example below makes the assumption that this service has users stored inside a sql dabase
and there are some helper methods attached to the account that abstract some of the data access for us.



```golang
import (
    "database/sql"

    "github.com/MovieStoreGuy/specify"
)

func AccountExists(db *sql.DB) specify.Condition[*AccountDetails] {
    return specify.ConditionFunc[*AccountDetails](func(account *AccountDetails) (bool, error){ 
        rows, err := db.query("valid query here...", account.Username())
        if err != nil {
            return false, err
        }
        return len(rows) == 1, nil
    })
}

func AccountSuspended() specify.Condition[*AccountDetails] {
    return specify.ConditionFunc[*AccountDetails](func(account *AccountDetails) (bool, error){ 
        return account.Deactivated() || time.Now().Sub(account.LastAcessed()) > 30 * 24 * time.Hour
    })
}

func AccountAdmin() specify.Condition[*AccountDetails] {
    return specify.ConditionFunc[*AccountDetails](func(account *AccountDetails) (bool, error){
        return account.Flags() & ADMIN_FLAG == ADMIN_FLAG, nil
    })
}


// IsValidAccount ensure that the account exists first and validates that the account is not suspended
// by chaining the above functions.
func IsValidAccount(db *sql.DB) specify.Condition[*AccountDetails] {
    return AccountExists(db).And(AccountSuspended().Not())
}

// IsValidAdmin extends the `IsValidAccount` by appending an 
// additional condition to verify admin flags are set.
func IsValidAdmin(db *sql.DB) specify.Condition[*AccountDetails] {
    return IsValidAccount(db).And(AdminAccount())
}
```

## Examples

Please see [Examples](./examples/) for potential ideas on how to use this.

<!-- References Link -->
[Specification Pattern]:https://en.wikipedia.org/wiki/Specification_pattern