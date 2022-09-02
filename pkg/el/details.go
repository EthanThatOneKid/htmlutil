package el

const (
	TAG_DETAILS = "details"
	TAG_SUMMARY = "summary"
)

func Details(tfs ...Tf) El {
	return New(TAG_DETAILS, tfs...)
}

func Summary(tfs ...Tf) El {
	return New(TAG_SUMMARY, tfs...)
}

func WithDetails(summary, details string) Tf {
	return func(e El) El {
		return e.Apply(WithChild(Details(WithSummary(summary), WithText(details))))
	}
}

func WithSummary(summary string) Tf {
	return func(e El) El {
		return e.Apply(WithChild(Summary(WithText(summary))))
	}
}
