# Система расчёта заработной платы сотрудников

Полностью объектная модель на Go: 61 структуры, >150 полей, >100 уникальных методов, 30+ ассоциаций между сущностями, 12 персональных исключений. Покрытие тестами `go test ./... -cover` >90%.

## Пакет apperrors (12 исключений)
- AccessError: Action, User; методы Error, Code, IsCritical.
- PayrollLockError: Period, Owner; методы Error, Code, RequiresEscalation.
- InsufficientBalanceError: Account, Needed, Actual; методы Error, Code, Shortfall.
- DuplicateEmployeeError: EmployeeID; методы Error, Code, IsSame.
- InvalidCurrencyError: Currency; методы Error, Code, IsEmpty.
- ScheduleConflictError: EmployeeID, Slot; методы Error, Code, IsSameSlot.
- PolicyViolationError: Policy, Actor; методы Error, Code, Involves.
- OvertimeLimitError: EmployeeID, Hours, Limit; методы Error, Code, Excess.
- UnauthenticatedActionError: Action; методы Error, Code, IsPublic.
- InconsistentStateError: Entity, Detail; методы Error, Code, IsEmpty.
- ApprovalMissingError: RequestID, Level; методы Error, Code, NeedsManager.
- DataIntegrityError: Field, Value; методы Error, Code, HasValue.

## Пакет hr
- Employee: ID, Name, Position, Contract, Salary, Department, Active; методы Activate, AssignPosition, UpdateSalary, AwardBonus; ассоциации Position, Contract.
- Position: Title, Level, BaseSalary; методы Promote, IsSenior.
- Contract: ID, EmployeeID, StartDate, EndDate, Rate, Type; методы IsActive, Extend.
- Permission: Name, Scope, Allowed; методы Grant, Revoke, CanAccess.
- AccessBadge: ID, EmployeeID, Active, LastUsed, Location; методы ActivateBadge, UseBadge, MatchesUser.
- Schedule: EmployeeID, WeeklyHours, Flexible, Shifts; методы AddShift, ToggleFlexible, HasShift.

## Пакет payroll
- SalaryBand: Level, Min, Max, Currency; методы Contains, Adjust.
- PayComponent: Name, Amount, Taxable, Recurring; методы ApplyTo, MakeNonRecurring, IsBonus.
- PayCalendar: Frequency, CycleDays, NextCutoff; методы IsValid, ShiftCutoff.
- Deduction: Name, Rate, Cap; методы Calculate, RateApplied.
- Bonus: Title, Amount, Recurring; методы Apply, StopRecurring, IsActive.
- OvertimePolicy: Name, RatePerHour, MaxHours, RequiresApproval; методы Calculate, EnableApproval, Allows.
- Paycheck: EmployeeID, Gross, Net, Deductions, Bonuses; методы AddDeduction, AddBonus, ComputeNet, TotalAdjustments.
- PayrollRun: ID, Period, Employees, Calendar, Components; методы AddEmployee, AddComponent, TotalPayroll, UsesCalendar; ассоциации hr.Employee, PayCalendar, PayComponent.

## Пакет finance
- Currency: Code, Symbol, Precision; методы Format, Equals.
- ExchangeRateTable: Base, Rates, LastUpdated; методы SetRate, Convert, HasRate; ассоциации Currency.
- BankAccount: Number, Owner, Balance, Currency, Active; методы Deposit, Withdraw, Activate, HasCurrency; ассоциации Currency.
- PaymentInstruction: ID, Amount, Currency, Receiver, Status; методы Approve, MarkPaid, IsReady; ассоциации Currency, BankAccount.
- LedgerEntry: ID, Debit, Credit, Account, Description; методы BalanceImpact, IsCredit; ассоциации BankAccount.
- TransactionBatch: ID, Entries, Total, Processed; методы AddEntry, ComputeTotal, Close, IsBalanced; ассоциации LedgerEntry.

## Пакет timeoff
- LeavePolicy: Name, AnnualLimit, CarryOver, RequiresApproval; методы EnableCarryOver, SetApproval, Allow.
- LeaveBalance: EmployeeID, Available, Taken, Policy; методы Accrue, Take, Remaining; ассоциации LeavePolicy.
- LeaveRequest: ID, EmployeeID, Days, Policy, Status; методы Approve, Reject, IsApproved; ассоциации LeavePolicy.
- LeaveApproval: Approver, Request, Comment, Approved; методы SignOff, IsFinal; ассоциации LeaveRequest.
- LeaveAccrual: EmployeeID, Rate, Accumulated; методы ApplyMonth, Reset, NextBalance; ассоциации LeaveBalance.

## Пакет benefits
- BenefitPlan: Name, Provider, Cost, EmployerContribution, Active; методы Activate, EmployerShare, EmployeeShare.
- BenefitEnrollment: EmployeeID, Plan, Status, StartDate; методы Cancel, Reactivate, IsActive; ассоциации BenefitPlan.
- InsuranceClaim: ID, Enrollment, Amount, Approved, Status, PayoutValue; методы Approve, Reject, Payout; ассоциации BenefitEnrollment.
- ReimbursementRequest: ID, EmployeeID, Amount, Category, Status; методы Approve, Reject, IsApproved.
- BenefitStatement: EmployeeID, Plans, TotalCost; методы AddPlan, ComputeTotal, PlanCount; ассоциации BenefitPlan.

## Пакет performance
- PerformanceGoal: ID, EmployeeID, Target, Achieved, Status; методы MarkProgress, Completion.
- PerformanceReview: ID, EmployeeID, Score, Reviewer, Feedback; методы UpdateScore, IsPassing, AddFeedback.
- PromotionCase: ID, Candidate, ProposedRole, Approved, Decision; методы Approve, Deny, Outcome; ассоциации hr.Employee, hr.Position.
- TrainingCourse: Name, Hours, Completed, Score; методы Complete, NeedsRetake, DurationHours.
- Certification: Name, Issuer, ValidUntil, Active; методы Activate, Expire, IsValid.

## Пакет compliance
- PolicyDocument: Name, Version, Active; методы Publish, Retire, IsActive.
- AuditLog: Actor, Action, Timestamp, Entries; методы AddEntry, EntryCount, LastEntry.
- AccessReview: ID, Reviewer, Findings, Approved; методы AddFinding, Approve, HasFindings.
- IncidentReport: ID, Severity, Description, Resolved, Resolution; методы Resolve, IsHighSeverity.
- RiskAssessment: ID, Risks, Threshold, Score; методы AddRisk, IsCritical.

## Пакет communication
- NotificationPreference: EmployeeID, Channel, Enabled, Frequency; методы Enable, Disable, UpdateFrequency.
- EmailTemplate: Name, Subject, Body, Active; методы Activate, Render, IsActive.
- MessageSender: SenderID, SentCount, Channel; методы Send, CanSend.
- AlertSubscription: ID, EmployeeID, Topic, Active; методы Subscribe, Unsubscribe, IsActive.
- ReminderSchedule: ID, Times, Enabled, LastSent; методы AddTime, Toggle, RecordSend, NextTime.

## Пакет reporting
- PayrollReport: Period, Total, GeneratedBy, Items; методы AddItem, ComputeTotal, HasData.
- HeadcountReport: DepartmentCounts, ActiveEmployees; методы AddDepartment, Total, DepartmentCount.
- ComplianceReport: Issues, RiskLevel, Reviewer; методы AddIssue, IsClean.
- BenefitsReport: EmployeeID, BenefitTotal, PlanCount; методы AddBenefit, AverageCost.

## Проверка
- Тесты: `go test ./...`
- Покрытие: `go test ./... -cover` (пакеты 90–96%+, общее >90%)
