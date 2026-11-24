A new git repo that focus on the domain design before creating stuff around.

Deliberately open to enable domain discussion

Mission: "In search of the Aggragate" 

Event chain
- BankAccountOpened
- AccountInBankAccountOpened
- DepotOpened
- SecurityBought
- DividendAnnounced
- ExDateReached
- PaymentReceived (!)

Finding the Right Boundaries
https://www.ompluscator.io/article/golang/practical-ddd-module/
https://milanjovanovic.tech/blog/aggregate-design-ddd


Ask these questions to find the boundaries:
- What must be consistent immediately? That's your Aggregate.
- What can be eventually consistent? That goes in a separate Aggregate with domain events.
- What data changes together? Group it.
- What's the smallest unit I can lock without breaking invariants? That's your ideal Aggregate size.

For example, an Order must be consistent with its OrderLineItems: the total must always reflect the items. But the Customer loyalty points can update eventually after the order completes.

Effective aggregate design (Vernon 2011)
https://www.dddcommunity.org/library/vernon_2011/
https://www.dddcommunity.org/wp-content/uploads/files/pdf_articles/Vernon_2011_1.pdf
https://www.dddcommunity.org/wp-content/uploads/files/pdf_articles/Vernon_2011_2.pdf
https://vimeo.com/33708293
https://www.dddcommunity.org/wp-content/uploads/files/pdf_articles/Vernon_2011_3.pdf
https://www.dddcommunity.org/library/vernon_2012