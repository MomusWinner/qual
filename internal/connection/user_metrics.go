package connection

import (
	"app/internal/domain/models"
	"app/internal/domain/repositories"
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type userRepoMetrics struct {
	userRepo repositories.UserRepository
	metrics  *metrics
}

type metrics struct {
	dbResponseTime *prometheus.HistogramVec
}

func newMetrics() *metrics {
	return &metrics{
		dbResponseTime: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "db_response_time_seconds",
				Help: "Database response time in seconds",
			},
			[]string{"method"},
		),
	}
}

func newUserRepositoryMetrics(userRepo repositories.UserRepository) *userRepoMetrics {
	return &userRepoMetrics{
		metrics:  newMetrics(),
		userRepo: userRepo,
	}
}

func (r *userRepoMetrics) Add(ctx context.Context, user models.User) (*models.User, error) {
	timer := prometheus.NewTimer(r.metrics.dbResponseTime.With(prometheus.Labels{"method": "Add"}))
	defer timer.ObserveDuration()
	return r.userRepo.Add(ctx, user)
}

func (r *userRepoMetrics) GetById(ctx context.Context, id int) (*models.User, error) {
	timer := prometheus.NewTimer(r.metrics.dbResponseTime.With(prometheus.Labels{"method": "GetById"}))
	defer timer.ObserveDuration()
	return r.userRepo.GetById(ctx, id)
}

func (r *userRepoMetrics) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	timer := prometheus.NewTimer(r.metrics.dbResponseTime.With(prometheus.Labels{"method": "GetById"}))
	defer timer.ObserveDuration()
	return r.userRepo.GetByEmail(ctx, email)
}

func (r *userRepoMetrics) GetAll(ctx context.Context) ([]models.User, error) {
	timer := prometheus.NewTimer(r.metrics.dbResponseTime.With(prometheus.Labels{"method": "GetAll"}))
	defer timer.ObserveDuration()
	return r.userRepo.GetAll(ctx)
}

func (r *userRepoMetrics) Update(ctx context.Context, user models.User) (*models.User, error) {
	timer := prometheus.NewTimer(r.metrics.dbResponseTime.With(prometheus.Labels{"method": "Update"}))
	defer timer.ObserveDuration()
	return r.userRepo.Update(ctx, user)
}

func (r *userRepoMetrics) Delete(ctx context.Context, id int) error {
	timer := prometheus.NewTimer(r.metrics.dbResponseTime.With(prometheus.Labels{"method": "Delete"}))
	defer timer.ObserveDuration()
	return r.userRepo.Delete(ctx, id)
}
